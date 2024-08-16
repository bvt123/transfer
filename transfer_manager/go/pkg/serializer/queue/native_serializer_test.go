package queue

import (
	"testing"

	"github.com/doublecloud/transfer/transfer_manager/go/pkg/abstract"
	server "github.com/doublecloud/transfer/transfer_manager/go/pkg/abstract/model"
	"github.com/stretchr/testify/require"
)

var nativeSerializerTestTypicalChangeItem *abstract.ChangeItem
var nativeSerializerTestTypicalChangeItems []abstract.ChangeItem

func init() {
	var testChangeItem = `{"id":601,"nextlsn":25051056,"commitTime":1643660670333075000,"txPosition":0,"kind":"insert","schema":"public","table":"basic_types15","part":"","columnnames":["id","val"],"columnvalues":[1,-8388605],"table_schema":[{"path":"","name":"id","type":"int32","key":true,"required":false,"original_type":"pg:integer"},{"path":"","name":"val","type":"int32","key":false,"required":false,"original_type":"pg:integer"}],"oldkeys":{},"tx_id":"","query":""}`
	nativeSerializerTestTypicalChangeItem, _ = abstract.UnmarshalChangeItem([]byte(testChangeItem))

	var testChangeItem0 = `{"id":1650228523,"nextlsn":2000000005662,"commitTime":1650228523000000000,"txPosition":0,"kind":"insert","schema":"source","table":"test_table","part":"","columnnames":["id","tinyint","tinyint_def","tinyint_u","tinyint_z","smallint","smallint_u","smallint_z","mediumint","mediumint_u","mediumint_z","int","int_u","int_z","bigint","bigint_u","bigint_z","bool","decimal_10_2","decimal_65_30","decimal_65_0","dec","numeric","fixed","float","float_z","float_53","real","double","double_precision","bit","bit_5","date","datetime","datetime_6","timestamp","timestamp_2","time","time_2","year","char","varchar","varchar_def","binary","varbinary","tinyblob","blob","mediumblob","longblob","tinytext","text","mediumtext","longtext","enum","set","json"],"columnvalues":[1,-128,-128,0,0,-32768,0,0,-8388608,0,0,-2147483648,0,0,-9223372036854775808,0,0,0,3.5,45.67,4567,4,3.5,4,10234568,3.5,3.5,3.5,-1.7976931348623157e+308,3.5,0,31,0,0,0,10801000000,10801000000,"-838:59:59","-838:59:59",1901,"0","","","","","","","","","","","","","1","1","{\"a\":\"b\"}"],"table_schema":[{"path":"","name":"id","type":"int32","key":true,"required":true,"original_type":"mysql:int(11)","original_type_params":null},{"path":"","name":"tinyint","type":"int8","key":false,"required":false,"original_type":"mysql:tinyint(4)","original_type_params":null},{"path":"","name":"tinyint_def","type":"int8","key":false,"required":false,"original_type":"mysql:tinyint(4)","original_type_params":null},{"path":"","name":"tinyint_u","type":"uint8","key":false,"required":false,"original_type":"mysql:tinyint(3) unsigned","original_type_params":null},{"path":"","name":"tinyint_z","type":"uint8","key":false,"required":false,"original_type":"mysql:tinyint(3) unsigned zerofill","original_type_params":null},{"path":"","name":"smallint","type":"int16","key":false,"required":false,"original_type":"mysql:smallint(6)","original_type_params":null},{"path":"","name":"smallint_u","type":"uint16","key":false,"required":false,"original_type":"mysql:smallint(5) unsigned","original_type_params":null},{"path":"","name":"smallint_z","type":"uint16","key":false,"required":false,"original_type":"mysql:smallint(5) unsigned zerofill","original_type_params":null},{"path":"","name":"mediumint","type":"int32","key":false,"required":false,"original_type":"mysql:mediumint(9)","original_type_params":null},{"path":"","name":"mediumint_u","type":"uint32","key":false,"required":false,"original_type":"mysql:mediumint(8) unsigned","original_type_params":null},{"path":"","name":"mediumint_z","type":"uint32","key":false,"required":false,"original_type":"mysql:mediumint(8) unsigned zerofill","original_type_params":null},{"path":"","name":"int","type":"int32","key":false,"required":false,"original_type":"mysql:int(11)","original_type_params":null},{"path":"","name":"int_u","type":"uint32","key":false,"required":false,"original_type":"mysql:int(10) unsigned","original_type_params":null},{"path":"","name":"int_z","type":"uint32","key":false,"required":false,"original_type":"mysql:int(10) unsigned zerofill","original_type_params":null},{"path":"","name":"bigint","type":"int64","key":false,"required":false,"original_type":"mysql:bigint(20)","original_type_params":null},{"path":"","name":"bigint_u","type":"uint64","key":false,"required":false,"original_type":"mysql:bigint(20) unsigned","original_type_params":null},{"path":"","name":"bigint_z","type":"uint64","key":false,"required":false,"original_type":"mysql:bigint(20) unsigned zerofill","original_type_params":null},{"path":"","name":"bool","type":"int8","key":false,"required":false,"original_type":"mysql:tinyint(1)","original_type_params":null},{"path":"","name":"decimal_10_2","type":"double","key":false,"required":false,"original_type":"mysql:decimal(10,2)","original_type_params":null},{"path":"","name":"decimal_65_30","type":"double","key":false,"required":false,"original_type":"mysql:decimal(65,30)","original_type_params":null},{"path":"","name":"decimal_65_0","type":"double","key":false,"required":false,"original_type":"mysql:decimal(65,0)","original_type_params":null},{"path":"","name":"dec","type":"double","key":false,"required":false,"original_type":"mysql:decimal(10,0)","original_type_params":null},{"path":"","name":"numeric","type":"double","key":false,"required":false,"original_type":"mysql:decimal(11,3)","original_type_params":null},{"path":"","name":"fixed","type":"double","key":false,"required":false,"original_type":"mysql:decimal(10,0)","original_type_params":null},{"path":"","name":"float","type":"double","key":false,"required":false,"original_type":"mysql:float(10,2)","original_type_params":null},{"path":"","name":"float_z","type":"double","key":false,"required":false,"original_type":"mysql:float(10,2) unsigned zerofill","original_type_params":null},{"path":"","name":"float_53","type":"double","key":false,"required":false,"original_type":"mysql:double","original_type_params":null},{"path":"","name":"real","type":"double","key":false,"required":false,"original_type":"mysql:double(10,2)","original_type_params":null},{"path":"","name":"double","type":"double","key":false,"required":false,"original_type":"mysql:double","original_type_params":null},{"path":"","name":"double_precision","type":"double","key":false,"required":false,"original_type":"mysql:double","original_type_params":null},{"path":"","name":"bit","type":"string","key":false,"required":false,"original_type":"mysql:bit(1)","original_type_params":null},{"path":"","name":"bit_5","type":"string","key":false,"required":false,"original_type":"mysql:bit(5)","original_type_params":null},{"path":"","name":"date","type":"date","key":false,"required":false,"original_type":"mysql:date","original_type_params":null},{"path":"","name":"datetime","type":"timestamp","key":false,"required":false,"original_type":"mysql:datetime","original_type_params":null},{"path":"","name":"datetime_6","type":"timestamp","key":false,"required":false,"original_type":"mysql:datetime(6)","original_type_params":null},{"path":"","name":"timestamp","type":"timestamp","key":false,"required":false,"original_type":"mysql:timestamp","original_type_params":null},{"path":"","name":"timestamp_2","type":"timestamp","key":false,"required":false,"original_type":"mysql:timestamp(2)","original_type_params":null},{"path":"","name":"time","type":"utf8","key":false,"required":false,"original_type":"mysql:time","original_type_params":null},{"path":"","name":"time_2","type":"utf8","key":false,"required":false,"original_type":"mysql:time(2)","original_type_params":null},{"path":"","name":"year","type":"utf8","key":false,"required":false,"original_type":"mysql:year(4)","original_type_params":null},{"path":"","name":"char","type":"utf8","key":false,"required":false,"original_type":"mysql:char(10)","original_type_params":null},{"path":"","name":"varchar","type":"utf8","key":false,"required":false,"original_type":"mysql:varchar(20)","original_type_params":null},{"path":"","name":"varchar_def","type":"utf8","key":false,"required":false,"original_type":"mysql:varchar(20)","original_type_params":null},{"path":"","name":"binary","type":"string","key":false,"required":false,"original_type":"mysql:binary(20)","original_type_params":null},{"path":"","name":"varbinary","type":"string","key":false,"required":false,"original_type":"mysql:varbinary(20)","original_type_params":null},{"path":"","name":"tinyblob","type":"string","key":false,"required":false,"original_type":"mysql:tinyblob","original_type_params":null},{"path":"","name":"blob","type":"string","key":false,"required":false,"original_type":"mysql:blob","original_type_params":null},{"path":"","name":"mediumblob","type":"string","key":false,"required":false,"original_type":"mysql:mediumblob","original_type_params":null},{"path":"","name":"longblob","type":"string","key":false,"required":false,"original_type":"mysql:longblob","original_type_params":null},{"path":"","name":"tinytext","type":"utf8","key":false,"required":false,"original_type":"mysql:tinytext","original_type_params":null},{"path":"","name":"text","type":"utf8","key":false,"required":false,"original_type":"mysql:text","original_type_params":null},{"path":"","name":"mediumtext","type":"utf8","key":false,"required":false,"original_type":"mysql:mediumtext","original_type_params":null},{"path":"","name":"longtext","type":"utf8","key":false,"required":false,"original_type":"mysql:longtext","original_type_params":null},{"path":"","name":"enum","type":"utf8","key":false,"required":false,"original_type":"mysql:enum('1','2','3')","original_type_params":null},{"path":"","name":"set","type":"utf8","key":false,"required":false,"original_type":"mysql:set('1','2','3')","original_type_params":null},{"path":"","name":"json","type":"any","key":false,"required":false,"original_type":"mysql:json","original_type_params":null}],"oldkeys":{},"tx_id":"","query":""}`
	var testChangeItem0Obj, _ = abstract.UnmarshalChangeItem([]byte(testChangeItem0))
	var testChangeItem1 = `{"id":1650228523,"nextlsn":2000000005662,"commitTime":1650228523000000000,"txPosition":1,"kind":"insert","schema":"source","table":"test_table","part":"","columnnames":["id","tinyint","tinyint_def","tinyint_u","tinyint_z","smallint","smallint_u","smallint_z","mediumint","mediumint_u","mediumint_z","int","int_u","int_z","bigint","bigint_u","bigint_z","bool","decimal_10_2","decimal_65_30","decimal_65_0","dec","numeric","fixed","float","float_z","float_53","real","double","double_precision","bit","bit_5","date","datetime","datetime_6","timestamp","timestamp_2","time","time_2","year","char","varchar","varchar_def","binary","varbinary","tinyblob","blob","mediumblob","longblob","tinytext","text","mediumtext","longtext","enum","set","json"],"columnvalues":[2,127,127,255,127,32767,65535,32767,8388607,16777215,8388607,2147483647,4294967295,2147483647,9223372036854774784,18446744073709549568,9223372036854774784,1,12345678.1,4567.89,456789,12345678,12345678.1,4,340282.34,12345678,12345678.1,12345678.1,-2.2250738585072014e-308,0,1,31,18646,4102444799000000,4102444799000000,2147483647000000,2147483647000000,"838:59:59","838:59:59",2155,"255",null,null,null,null,null,null,null,null,null,null,null,null,"3","3","{\"a\":\"b\",\"c\":1,\"d\":{},\"e\":[]}"],"table_schema":[{"path":"","name":"id","type":"int32","key":true,"required":true,"original_type":"mysql:int(11)","original_type_params":null},{"path":"","name":"tinyint","type":"int8","key":false,"required":false,"original_type":"mysql:tinyint(4)","original_type_params":null},{"path":"","name":"tinyint_def","type":"int8","key":false,"required":false,"original_type":"mysql:tinyint(4)","original_type_params":null},{"path":"","name":"tinyint_u","type":"uint8","key":false,"required":false,"original_type":"mysql:tinyint(3) unsigned","original_type_params":null},{"path":"","name":"tinyint_z","type":"uint8","key":false,"required":false,"original_type":"mysql:tinyint(3) unsigned zerofill","original_type_params":null},{"path":"","name":"smallint","type":"int16","key":false,"required":false,"original_type":"mysql:smallint(6)","original_type_params":null},{"path":"","name":"smallint_u","type":"uint16","key":false,"required":false,"original_type":"mysql:smallint(5) unsigned","original_type_params":null},{"path":"","name":"smallint_z","type":"uint16","key":false,"required":false,"original_type":"mysql:smallint(5) unsigned zerofill","original_type_params":null},{"path":"","name":"mediumint","type":"int32","key":false,"required":false,"original_type":"mysql:mediumint(9)","original_type_params":null},{"path":"","name":"mediumint_u","type":"uint32","key":false,"required":false,"original_type":"mysql:mediumint(8) unsigned","original_type_params":null},{"path":"","name":"mediumint_z","type":"uint32","key":false,"required":false,"original_type":"mysql:mediumint(8) unsigned zerofill","original_type_params":null},{"path":"","name":"int","type":"int32","key":false,"required":false,"original_type":"mysql:int(11)","original_type_params":null},{"path":"","name":"int_u","type":"uint32","key":false,"required":false,"original_type":"mysql:int(10) unsigned","original_type_params":null},{"path":"","name":"int_z","type":"uint32","key":false,"required":false,"original_type":"mysql:int(10) unsigned zerofill","original_type_params":null},{"path":"","name":"bigint","type":"int64","key":false,"required":false,"original_type":"mysql:bigint(20)","original_type_params":null},{"path":"","name":"bigint_u","type":"uint64","key":false,"required":false,"original_type":"mysql:bigint(20) unsigned","original_type_params":null},{"path":"","name":"bigint_z","type":"uint64","key":false,"required":false,"original_type":"mysql:bigint(20) unsigned zerofill","original_type_params":null},{"path":"","name":"bool","type":"int8","key":false,"required":false,"original_type":"mysql:tinyint(1)","original_type_params":null},{"path":"","name":"decimal_10_2","type":"double","key":false,"required":false,"original_type":"mysql:decimal(10,2)","original_type_params":null},{"path":"","name":"decimal_65_30","type":"double","key":false,"required":false,"original_type":"mysql:decimal(65,30)","original_type_params":null},{"path":"","name":"decimal_65_0","type":"double","key":false,"required":false,"original_type":"mysql:decimal(65,0)","original_type_params":null},{"path":"","name":"dec","type":"double","key":false,"required":false,"original_type":"mysql:decimal(10,0)","original_type_params":null},{"path":"","name":"numeric","type":"double","key":false,"required":false,"original_type":"mysql:decimal(11,3)","original_type_params":null},{"path":"","name":"fixed","type":"double","key":false,"required":false,"original_type":"mysql:decimal(10,0)","original_type_params":null},{"path":"","name":"float","type":"double","key":false,"required":false,"original_type":"mysql:float(10,2)","original_type_params":null},{"path":"","name":"float_z","type":"double","key":false,"required":false,"original_type":"mysql:float(10,2) unsigned zerofill","original_type_params":null},{"path":"","name":"float_53","type":"double","key":false,"required":false,"original_type":"mysql:double","original_type_params":null},{"path":"","name":"real","type":"double","key":false,"required":false,"original_type":"mysql:double(10,2)","original_type_params":null},{"path":"","name":"double","type":"double","key":false,"required":false,"original_type":"mysql:double","original_type_params":null},{"path":"","name":"double_precision","type":"double","key":false,"required":false,"original_type":"mysql:double","original_type_params":null},{"path":"","name":"bit","type":"string","key":false,"required":false,"original_type":"mysql:bit(1)","original_type_params":null},{"path":"","name":"bit_5","type":"string","key":false,"required":false,"original_type":"mysql:bit(5)","original_type_params":null},{"path":"","name":"date","type":"date","key":false,"required":false,"original_type":"mysql:date","original_type_params":null},{"path":"","name":"datetime","type":"timestamp","key":false,"required":false,"original_type":"mysql:datetime","original_type_params":null},{"path":"","name":"datetime_6","type":"timestamp","key":false,"required":false,"original_type":"mysql:datetime(6)","original_type_params":null},{"path":"","name":"timestamp","type":"timestamp","key":false,"required":false,"original_type":"mysql:timestamp","original_type_params":null},{"path":"","name":"timestamp_2","type":"timestamp","key":false,"required":false,"original_type":"mysql:timestamp(2)","original_type_params":null},{"path":"","name":"time","type":"utf8","key":false,"required":false,"original_type":"mysql:time","original_type_params":null},{"path":"","name":"time_2","type":"utf8","key":false,"required":false,"original_type":"mysql:time(2)","original_type_params":null},{"path":"","name":"year","type":"utf8","key":false,"required":false,"original_type":"mysql:year(4)","original_type_params":null},{"path":"","name":"char","type":"utf8","key":false,"required":false,"original_type":"mysql:char(10)","original_type_params":null},{"path":"","name":"varchar","type":"utf8","key":false,"required":false,"original_type":"mysql:varchar(20)","original_type_params":null},{"path":"","name":"varchar_def","type":"utf8","key":false,"required":false,"original_type":"mysql:varchar(20)","original_type_params":null},{"path":"","name":"binary","type":"string","key":false,"required":false,"original_type":"mysql:binary(20)","original_type_params":null},{"path":"","name":"varbinary","type":"string","key":false,"required":false,"original_type":"mysql:varbinary(20)","original_type_params":null},{"path":"","name":"tinyblob","type":"string","key":false,"required":false,"original_type":"mysql:tinyblob","original_type_params":null},{"path":"","name":"blob","type":"string","key":false,"required":false,"original_type":"mysql:blob","original_type_params":null},{"path":"","name":"mediumblob","type":"string","key":false,"required":false,"original_type":"mysql:mediumblob","original_type_params":null},{"path":"","name":"longblob","type":"string","key":false,"required":false,"original_type":"mysql:longblob","original_type_params":null},{"path":"","name":"tinytext","type":"utf8","key":false,"required":false,"original_type":"mysql:tinytext","original_type_params":null},{"path":"","name":"text","type":"utf8","key":false,"required":false,"original_type":"mysql:text","original_type_params":null},{"path":"","name":"mediumtext","type":"utf8","key":false,"required":false,"original_type":"mysql:mediumtext","original_type_params":null},{"path":"","name":"longtext","type":"utf8","key":false,"required":false,"original_type":"mysql:longtext","original_type_params":null},{"path":"","name":"enum","type":"utf8","key":false,"required":false,"original_type":"mysql:enum('1','2','3')","original_type_params":null},{"path":"","name":"set","type":"utf8","key":false,"required":false,"original_type":"mysql:set('1','2','3')","original_type_params":null},{"path":"","name":"json","type":"any","key":false,"required":false,"original_type":"mysql:json","original_type_params":null}],"oldkeys":{},"tx_id":"","query":""}`
	var testChangeItem1Obj, _ = abstract.UnmarshalChangeItem([]byte(testChangeItem1))
	var testChangeItem2 = `{"id":5693,"nextlsn":2000000005693,"commitTime":1650228524794705194,"txPosition":0,"kind":"DDL","schema":"","table":"","part":"","columnnames":["query"],"columnvalues":["-- transaction  done"],"table_schema":[{"path":"","name":"","type":"","key":false,"required":false,"original_type":"","original_type_params":null}],"oldkeys":{},"tx_id":"5693","query":""}`
	var testChangeItem2Obj, _ = abstract.UnmarshalChangeItem([]byte(testChangeItem2))
	nativeSerializerTestTypicalChangeItems = []abstract.ChangeItem{*testChangeItem0Obj, *testChangeItem1Obj, *testChangeItem2Obj}
}

func TestNativeSerializerEmptyInput(t *testing.T) {
	batcher := server.Batching{
		Enabled:        false,
		Interval:       0,
		MaxChangeItems: 0,
		MaxMessageSize: 0,
	}
	nativeSerializer, err := NewNativeSerializer(batcher, true)
	require.NoError(t, err)

	batches, err := nativeSerializer.Serialize([]abstract.ChangeItem{})
	require.NoError(t, err)
	require.Len(t, batches, 0)
}

func TestNativeSerializerTopicName(t *testing.T) {
	batcher := server.Batching{
		Enabled:        false,
		Interval:       0,
		MaxChangeItems: 0,
		MaxMessageSize: 0,
	}

	changeItem0 := *nativeSerializerTestTypicalChangeItem
	changeItem0.Table = "table0"
	changeItem1 := *nativeSerializerTestTypicalChangeItem
	changeItem1.Table = "table1"

	//-------------------------
	// saveTxOrder: false

	nativeSerializer, err := NewNativeSerializer(batcher, false)
	require.NoError(t, err)

	batches, err := nativeSerializer.Serialize([]abstract.ChangeItem{changeItem0, changeItem1})
	require.NoError(t, err)
	require.True(t, len(batches) > 0)

	var ok bool
	var v []SerializedMessage

	v, ok = batches[abstract.TablePartID{TableID: abstract.TableID{Namespace: "public", Name: "table0"}, PartID: ""}]
	require.True(t, ok)
	require.Equal(t, `public_table0`, string(v[0].Key))
	require.Equal(t, `{"id":601,"nextlsn":25051056,"commitTime":1643660670333075000,"txPosition":0,"kind":"insert","schema":"public","table":"table0","part":"","columnnames":["id","val"],"columnvalues":[1,-8388605],"table_schema":[{"table_schema":"","table_name":"","path":"","name":"id","type":"int32","key":true,"fake_key":false,"required":false,"expression":"","original_type":"pg:integer"},{"table_schema":"","table_name":"","path":"","name":"val","type":"int32","key":false,"fake_key":false,"required":false,"expression":"","original_type":"pg:integer"}],"oldkeys":{},"tx_id":"","query":""}`, string(v[0].Value))

	v, ok = batches[abstract.TablePartID{TableID: abstract.TableID{Namespace: "public", Name: "table1"}, PartID: ""}]
	require.True(t, ok)
	require.Equal(t, `public_table1`, string(v[0].Key))
	require.Equal(t, `{"id":601,"nextlsn":25051056,"commitTime":1643660670333075000,"txPosition":0,"kind":"insert","schema":"public","table":"table1","part":"","columnnames":["id","val"],"columnvalues":[1,-8388605],"table_schema":[{"table_schema":"","table_name":"","path":"","name":"id","type":"int32","key":true,"fake_key":false,"required":false,"expression":"","original_type":"pg:integer"},{"table_schema":"","table_name":"","path":"","name":"val","type":"int32","key":false,"fake_key":false,"required":false,"expression":"","original_type":"pg:integer"}],"oldkeys":{},"tx_id":"","query":""}`, string(v[0].Value))

	//------------------------
	// saveTxOrder: true

	nativeSerializerTx, err := NewNativeSerializer(batcher, true)
	require.NoError(t, err)
	batchesTx, err := nativeSerializerTx.Serialize([]abstract.ChangeItem{changeItem0, changeItem1})
	require.NoError(t, err)
	require.Len(t, batchesTx, 1)
	for k, v := range batchesTx {
		require.Equal(t, abstract.TablePartID{TableID: abstract.TableID{Namespace: "", Name: ""}, PartID: ""}, k)
		require.Len(t, v, 2)
	}
}

func TestBatchNativeSerializerTopicName(t *testing.T) {
	batcher := server.Batching{
		Enabled:        false,
		Interval:       0,
		MaxChangeItems: 0,
		MaxMessageSize: 0,
	}
	nativeSerializer, err := NewNativeSerializer(batcher, true)
	require.NoError(t, err)

	msgs, err := nativeSerializer.Serialize(nativeSerializerTestTypicalChangeItems)
	require.NoError(t, err)
	require.Len(t, msgs, 1)
}
