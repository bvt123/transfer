package postgres

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/doublecloud/transfer/library/go/core/metrics"
	"github.com/doublecloud/transfer/library/go/core/metrics/solomon"
	"github.com/doublecloud/transfer/library/go/core/xerrors"
	"github.com/doublecloud/transfer/transfer_manager/go/internal/logger"
	"github.com/doublecloud/transfer/transfer_manager/go/pkg/dbaas"
	"github.com/stretchr/testify/require"
)

// docker-compose is used from:
//     - https://habr.com/ru/articles/754168/
// aka
//     - https://github.com/mfvanek/useful-sql-scripts/blob/master/running_pg_cluster_in_docker/1.%20Hello%20Postgres%20cluster/docker-compose.yml

func checkHostIsPrimary(cfg *PgStorageParams, host string) (bool, error) {
	resultHost, resultPort, err := dbaas.ResolveHostPortWithOverride(host, uint16(cfg.Port))
	if err != nil {
		return false, xerrors.Errorf("unable to extract host/port, host:%s, port:%d", host, cfg.Port)
	}

	conn, err := MakeConnPoolFromHostPort(cfg, resultHost, int(resultPort), logger.Log)
	if err != nil {
		return false, xerrors.Errorf("unable to make conn pool, err: %w", err)
	}
	defer conn.Close()

	query := `select case when pg_is_in_recovery() then 'secondary' else 'primary' end as host_status;`
	var status string
	err = conn.QueryRow(context.TODO(), query).Scan(&status)
	if err != nil {
		return false, xerrors.Errorf("unable to select primary/secondary, err: %w", err)
	}

	if status == "primary" {
		return true, nil
	} else if status == "secondary" {
		return false, nil
	} else {
		return false, xerrors.Errorf("unknown status: %s", status)
	}
}

func getMaster(cfg *PgStorageParams, hosts []string) (string, error) {
	for _, host := range hosts {
		isPrimary, _ := checkHostIsPrimary(cfg, host)
		if isPrimary {
			fmt.Println("getMaster(): determined master (ignore error): ", host)
			return host, nil
		}
		fmt.Printf("getMaster(): hosts: %s, isPrimary:%v\n", host, isPrimary)
	}
	for _, host := range hosts {
		isPrimary, err := checkHostIsPrimary(cfg, host)
		if err != nil {
			fmt.Println("getMaster(): checkHostIsPrimary returned an error: ", err.Error())
			return "", xerrors.New("blablabla")
		}
		if isPrimary {
			fmt.Println("determined master (handle error): ", host)
			return host, nil
		}
	}
	fmt.Println("getMaster(): no one host is primary - return empty string & error")
	return "", xerrors.New("blablabla2")
}

func waitMaster(cfg *PgStorageParams, hosts []string, duration time.Duration) (string, error) {
	startTime := time.Now()
	for {
		master, err := getMaster(cfg, hosts)
		if err != nil {
			realDuration := time.Since(startTime)
			if realDuration > duration {
				return "", xerrors.New("realDuration > duration")
			}
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("return master: ", master)
		return master, nil
	}
}

// checkTwoPostgresIsRunning
// returns 'true' if needed to check one more time
func checkTwoPostgresIsRunning(t *testing.T) bool {
	cmd := `docker ps | grep '\bUp\b.*\bpostgres_\d\b' | wc -l | awk '{print $1}'`
	out, err := exec.Command("bash", "-c", cmd).Output()
	require.NoError(t, err)
	num, err := strconv.Atoi(strings.Trim(string(out), "\n"))
	require.NoError(t, err)
	if num == 2 {
		fmt.Println("there are 2 upraised postgres")
		return false
	} else {
		fmt.Println("there are NOT 2 upraised postgres - will start postgres_1/postgres_2 and try again")
		var err error
		_, err = exec.Command("bash", "-c", `docker start postgres_1`).Output()
		require.NoError(t, err)
		_, err = exec.Command("bash", "-c", `docker start postgres_2`).Output()
		require.NoError(t, err)
		time.Sleep(2 * time.Second)
		return true
	}
}

// checkExactlyOnePostgresIsPrimary
// returns 'true' if there are two primary
func checkExactlyOnePostgresIsPrimary(t *testing.T) bool {
	cmd1 := `docker exec postgres_1 psql -t -c "select case when pg_is_in_recovery() then 'secondary' else 'primary' end as host_status;" "dbname=habrdb user=habrpguser password=pgpwd4habr" | awk '{print $1}' | grep -v "^$"`
	out1Arr, err := exec.Command("bash", "-c", cmd1).Output()
	require.NoError(t, err)
	out1 := strings.Trim(string(out1Arr), "\n")
	fmt.Println("postgres_1 status:", out1)

	cmd2 := `docker exec postgres_2 psql -t -c "select case when pg_is_in_recovery() then 'secondary' else 'primary' end as host_status;" "dbname=habrdb user=habrpguser password=pgpwd4habr" | awk '{print $1}' | grep -v "^$"`
	out2Arr, err := exec.Command("bash", "-c", cmd2).Output()
	require.NoError(t, err)
	out2 := strings.Trim(string(out2Arr), "\n")
	fmt.Println("postgres_2 status:", out2)

	if out1 == "primary" && out2 == "primary" {
		fmt.Println("both nodes are primary")
		return false
	}
	if out1 == "primary" && out2 != "primary" {
		return true
	}
	if out1 != "primary" && out2 == "primary" {
		return true
	}
	panic("!")
}

func restartPostgres1(t *testing.T) {
	var err error
	_, err = exec.Command("bash", "-c", `docker stop postgres_1`).Output()
	require.NoError(t, err)
	_, err = exec.Command("bash", "-c", `docker start postgres_1`).Output()
	require.NoError(t, err)
	time.Sleep(2 * time.Second)
}

func ensureStorageCreatingOnMaster(t *testing.T, cfg *PgStorageParams, hosts []string, expectedMasterHost string) {
	cfg.AllHosts = hosts
	storage, err := NewStorage(cfg)
	require.NoError(t, err)
	defer storage.Close()
	storageHost := fmt.Sprintf("%s:%d", storage.Conn.Config().ConnConfig.Config.Host, storage.Conn.Config().ConnConfig.Config.Port)
	require.Equal(t, expectedMasterHost, storageHost)
}

func emptyRegistry() metrics.Registry {
	return solomon.NewRegistry(nil).WithTags(map[string]string{"ts": time.Now().String()})
}

func ensureSinkCreatingOnMaster(t *testing.T, cfg *PgDestination, hosts []string, expectedMasterHost string) {
	cfg.Hosts = hosts
	currSink, err := NewSink(logger.Log, "dtt", cfg.ToSinkParams(), emptyRegistry())
	currSinkUnwrapped := currSink.(*sink)
	require.NoError(t, err)
	defer currSink.Close()
	storageHost := fmt.Sprintf("%s:%d", currSinkUnwrapped.conn.Config().ConnConfig.Config.Host, currSinkUnwrapped.conn.Config().ConnConfig.Config.Port)
	require.Equal(t, expectedMasterHost, storageHost)
}

func shutdownContainer(t *testing.T, cfg *PgStorageParams, host string) {
	_, resultPort, err := dbaas.ResolveHostPortWithOverride(host, uint16(cfg.Port))
	require.NoError(t, err)

	substr := fmt.Sprintf("0.0.0.0:%d->", resultPort)

	cmd := fmt.Sprintf(`docker ps | grep "%s" | awk '{print $1}' | xargs docker kill`, substr)
	_, err = exec.Command("bash", "-c", cmd).Output()
	require.NoError(t, err)
}

func prepare(t *testing.T) {
	for {
		fmt.Println(strings.Repeat("-", 25))
		if checkTwoPostgresIsRunning(t) {
			continue
		}
		if checkExactlyOnePostgresIsPrimary(t) {
			break
		}
		fmt.Println("will restart postgres_1")
		restartPostgres1(t)
		fmt.Println("postgres_1 restarted")
	}
}

func TestStorage(t *testing.T) {
	t.Skip()
	prepare(t)

	// action

	hosts := []string{
		"localhost:6432",
		"localhost:6433",
	}

	dst := &PgSource{
		Hosts:    hosts,
		Port:     9999, // should be never used
		Database: "habrdb",
		User:     "habrpguser",
		Password: "pgpwd4habr",
	}
	cfg := dst.ToStorageParams(nil)

	masterHost, err := getMaster(cfg, hosts)
	require.NoError(t, err)
	ensureStorageCreatingOnMaster(t, cfg, []string{hosts[0], hosts[1]}, masterHost) // direct order
	ensureStorageCreatingOnMaster(t, cfg, []string{hosts[1], hosts[0]}, masterHost) // reverse order

	fmt.Printf("will shutdown current master: %s\n", masterHost)

	shutdownContainer(t, cfg, masterHost)

	fmt.Println("current master gunned down, will wait when new node became master")

	masterHost2, err := waitMaster(cfg, hosts, time.Second*60)
	require.NoError(t, err)
	require.NotEqual(t, masterHost, masterHost2)

	fmt.Printf("got new master: %s, will check storage\n", masterHost2)

	ensureStorageCreatingOnMaster(t, cfg, []string{hosts[0], hosts[1]}, masterHost2) // direct order
	ensureStorageCreatingOnMaster(t, cfg, []string{hosts[1], hosts[0]}, masterHost2) // reverse order

	fmt.Println("everything is ok!")
}

func TestSink(t *testing.T) {
	t.Skip()
	prepare(t)

	// action

	hosts := []string{
		"localhost:6432",
		"localhost:6433",
	}

	dst := &PgDestination{
		Hosts:    hosts,
		Port:     9999, // should be never used
		Database: "habrdb",
		User:     "habrpguser",
		Password: "pgpwd4habr",
	}
	cfgAsStorage := dst.ToStorageParams()

	masterHost, err := getMaster(cfgAsStorage, hosts)
	require.NoError(t, err)
	ensureSinkCreatingOnMaster(t, dst, []string{hosts[0], hosts[1]}, masterHost) // direct order
	ensureSinkCreatingOnMaster(t, dst, []string{hosts[1], hosts[0]}, masterHost) // reverse order

	fmt.Printf("will shutdown current master: %s\n", masterHost)

	shutdownContainer(t, cfgAsStorage, masterHost)

	fmt.Println("current master gunned down, will wait when new node became master")

	masterHost2, err := waitMaster(cfgAsStorage, hosts, time.Second*60)
	require.NoError(t, err)
	require.NotEqual(t, masterHost, masterHost2)

	fmt.Printf("got new master: %s, will check storage\n", masterHost2)

	ensureSinkCreatingOnMaster(t, dst, []string{hosts[0], hosts[1]}, masterHost2) // direct order
	ensureSinkCreatingOnMaster(t, dst, []string{hosts[1], hosts[0]}, masterHost2) // reverse order

	fmt.Println("everything is ok!")
}
