package kafka

import (
	"github.com/doublecloud/transfer/library/go/core/xerrors"
	"github.com/doublecloud/transfer/transfer_manager/go/pkg/abstract"
	server "github.com/doublecloud/transfer/transfer_manager/go/pkg/abstract/model"
	debeziumparameters "github.com/doublecloud/transfer/transfer_manager/go/pkg/debezium/parameters"
	"github.com/doublecloud/transfer/transfer_manager/go/pkg/middlewares/async/bufferer"
)

type KafkaDestination struct {
	Connection       *KafkaConnectionOptions
	Auth             *KafkaAuth
	SecurityGroupIDs []string

	// The setting from segmentio/kafka-go Writer.
	// Tunes max length of one message (see usages of BatchBytes in kafka-go)
	// Msg size: len(key)+len(val)+14
	// By default is 0 - then kafka-go set it into 1048576.
	// When set it to not default - remember than managed kafka (server-side) has default max.message.bytes == 1048588
	BatchBytes          int64
	ParralelWriterCount int

	Topic       string // full-name version
	TopicPrefix string

	AddSystemTables bool // private options - to not skip consumer_keeper & other system tables
	SaveTxOrder     bool

	// for now, 'FormatSettings' is private option - it's WithDefaults(): SerializationFormatAuto - 'Mirror' for queues, 'Debezium' for the rest
	FormatSettings server.SerializationFormat

	TopicConfigEntries []TopicConfigEntry

	// Compression which compression mechanism use for writer, default - None
	Compression Encoding
}

var _ server.Destination = (*KafkaDestination)(nil)

type TopicConfigEntry struct {
	ConfigName, ConfigValue string
}

func (d *KafkaDestination) MDBClusterID() string {
	if d.Connection != nil {
		return d.Connection.ClusterID
	}
	return ""
}

func (d *KafkaDestination) WithDefaults() {
	if d.Connection == nil {
		d.Connection = &KafkaConnectionOptions{
			ClusterID:    "",
			TLS:          "",
			TLSFile:      "",
			Brokers:      nil,
			SubNetworkID: "",
		}
	}
	if d.Auth == nil {
		d.Auth = &KafkaAuth{
			Enabled:   true,
			Mechanism: "SHA-512",
			User:      "",
			Password:  "",
		}
	}
	if d.FormatSettings.Name == "" {
		d.FormatSettings.Name = server.SerializationFormatAuto
	}
	if d.FormatSettings.Settings == nil {
		d.FormatSettings.Settings = make(map[string]string)
	}
	if d.FormatSettings.BatchingSettings == nil {
		d.FormatSettings.BatchingSettings = &server.Batching{
			Enabled:        false,
			Interval:       0,
			MaxChangeItems: 0,
			MaxMessageSize: 0,
		}
	}
	if d.ParralelWriterCount == 0 {
		d.ParralelWriterCount = 10
	}
}

func (d *KafkaDestination) CleanupMode() server.CleanupType {
	return server.DisabledCleanup
}

func (d *KafkaDestination) Transformer() map[string]string {
	return nil
}

func (KafkaDestination) IsDestination() {}

func (d *KafkaDestination) GetProviderType() abstract.ProviderType {
	return ProviderType
}

func (d *KafkaDestination) Validate() error {
	if d.TopicPrefix != "" && d.SaveTxOrder {
		return xerrors.Errorf("option 'SaveTxOrder'=true is incompatible with 'TopicPrefix'. Use either full topic name or turn off 'SaveTxOrder'.")
	}
	return nil
}

func (d *KafkaDestination) Compatible(src server.Source, transferType abstract.TransferType) error {
	return sourceCompatible(src, transferType, d.FormatSettings.Name)
}

func (d *KafkaDestination) Serializer() (server.SerializationFormat, bool) {
	formatSettings := d.FormatSettings
	formatSettings.Settings = debeziumparameters.GetDefaultParameters(formatSettings.Settings)
	return formatSettings, d.SaveTxOrder
}

func (d *KafkaDestination) BuffererConfig() bufferer.BuffererConfig {
	return bufferer.BuffererConfig{
		TriggingCount:    d.FormatSettings.BatchingSettings.MaxChangeItems,
		TriggingSize:     uint64(d.FormatSettings.BatchingSettings.MaxMessageSize),
		TriggingInterval: d.FormatSettings.BatchingSettings.Interval,
	}
}

var _ server.HostResolver = (*KafkaDestination)(nil)

func (d *KafkaDestination) HostsNames() ([]string, error) {
	if d.Connection != nil && d.Connection.ClusterID != "" {
		return nil, nil
	}
	return ResolveOnPremBrokers(d.Connection, d.Auth)
}
