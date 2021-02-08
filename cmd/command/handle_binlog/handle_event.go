package handle_binlog

import (
	"encoding/json"
	"fmt"
	"os"
	"owen2020/app/apputil"
	"owen2020/app/models"
	"owen2020/conn"
	"reflect"
	"strconv"
	"strings"
	"time"
	"xorm.io/core"

	"github.com/siddontang/go-mysql/replication"

	"github.com/google/go-cmp/cmp"
)

//Filter 定义过滤
var Filter string = ""

//Cfg mysql 配置
var Cfg replication.BinlogSyncerConfig

//TableColumnIdentify 表字段ID对应字段名
type TableColumnIdentify map[int]string

//DBTable  数据库.数据表map类型
type DBTable map[string]TableColumnIdentify

// DBTables 数据库.数据表map配置实例
var DBTables DBTable = DBTable{}

// HandleEvent mysql binlog event 处理
func HandleEvent(e *replication.BinlogEvent) {
	switch e.Header.EventType {
	// case replication.UNKNOWN_EVENT:
	// case replication.START_EVENT_V3:
	// 	return "StartEventV3"
	case replication.QUERY_EVENT: // 数据表结构变更和新增表，如 query:ALTER TABLE `codeper`.`user` ADD COLUMN `test` varchar(255) NULL AFTER `age`
		// apputil.PrettyPrint(e)
		handleQueryEvent(e)
	// 	return "QueryEvent"
	// case replication.STOP_EVENT:
	// 	return "StopEvent"
	case replication.ROTATE_EVENT: // 上一个mysql binlog file结束标识
	// 	return "RotateEvent"
	// case replication.INTVAR_EVENT:
	// 	return "IntVarEvent"
	// case replication.LOAD_EVENT:
	// 	return "LoadEvent"
	// case replication.SLAVE_EVENT:
	// 	return "SlaveEvent"
	// case replication.CREATE_FILE_EVENT:
	// 	return "CreateFileEvent"
	// case replication.APPEND_BLOCK_EVENT:
	// 	return "AppendBlockEvent"
	// case replication.EXEC_LOAD_EVENT:
	// 	return "ExecLoadEvent"
	// case replication.DELETE_FILE_EVENT:
	// 	return "DeleteFileEvent"
	// case replication.NEW_LOAD_EVENT:
	// 	return "NewLoadEvent"
	// case replication.RAND_EVENT:
	// 	return "RandEvent"
	// case replication.USER_VAR_EVENT:
	// 	return "UserVarEvent"
	case replication.FORMAT_DESCRIPTION_EVENT: // 新mysql binlog file 开始标识
	// 	return "FormatDescriptionEvent"
	// case replication.XID_EVENT:
	// 	return "XIDEvent"
	// case replication.BEGIN_LOAD_QUERY_EVENT:
	// 	return "BeginLoadQueryEvent"
	// case replication.EXECUTE_LOAD_QUERY_EVENT:
	// 	return "ExectueLoadQueryEvent"
	case replication.TABLE_MAP_EVENT: // 有insert, update, delete)操作时，第来一条TABLE_MAP_EVENT 说明当下表结构
		// apputil.PrettyPrint(e)
		// ev, _ := e.Event.(*replication.TableMapEvent)
		// e.Event.(replication.TableMapEvent).Table
		// FlushTableIdentifierNameMap(string(ev.Schema), string(ev.Table))
	// 	return "TableMapEvent"
	case replication.WRITE_ROWS_EVENTv0:
	// 	return "WriteRowsEventV0"
	case replication.UPDATE_ROWS_EVENTv0:
		fmt.Println("update rows event v0")
		// apputil.PrettyPrint(e)
	// 	return "UpdateRowsEventV0"
	case replication.DELETE_ROWS_EVENTv0:
	// 	return "DeleteRowsEventV0"
	case replication.WRITE_ROWS_EVENTv1:
		// apputil.PrettyPrint(e)
		handleWriteRowsEventV1(e)
	// 	return "WriteRowsEventV1"
	case replication.UPDATE_ROWS_EVENTv1:
		apputil.PrettyPrint(e)
		handleUpdateEventV1(e)
	// 	return "UpdateRowsEventV1"
	case replication.DELETE_ROWS_EVENTv1:
		// apputil.PrettyPrint(e)
		handleDeleteRowsEventV1(e)
	// 	return "DeleteRowsEventV1"
	// case replication.INCIDENT_EVENT:
	// 	return "IncidentEvent"
	// case replication.HEARTBEAT_EVENT:
	// 	return "HeartbeatEvent"
	// case replication.IGNORABLE_EVENT:
	// 	return "IgnorableEvent"
	// case replication.ROWS_QUERY_EVENT:
	// 	return "RowsQueryEvent"
	case replication.WRITE_ROWS_EVENTv2:
	case replication.UPDATE_ROWS_EVENTv2:
		fmt.Println("update rows event v2")
		apputil.PrettyPrint(e)
	case replication.DELETE_ROWS_EVENTv2:

	case replication.GTID_EVENT: // 新GTID开始标识，它前面必然有一个PREVIOUS_GTIDS_EVENT事件
	// 	return "GTIDEvent"
	case replication.ANONYMOUS_GTID_EVENT:
	// 	return "AnonymousGTIDEvent"
	case replication.PREVIOUS_GTIDS_EVENT: //老GTID结束标识
	// 	return "PreviousGTIDsEvent"
	// case replication.MARIADB_ANNOTATE_ROWS_EVENT:
	// 	return "MariadbAnnotateRowsEvent"
	// case replication.MARIADB_BINLOG_CHECKPOINT_EVENT:
	// 	return "MariadbBinLogCheckPointEvent"
	// case replication.MARIADB_GTID_EVENT:
	// 	return "MariadbGTIDEvent"
	// case replication.MARIADB_GTID_LIST_EVENT:
	// 	return "MariadbGTIDListEvent"
	case replication.TRANSACTION_CONTEXT_EVENT:
	// 	return "TransactionContextEvent"
	// case replication.VIEW_CHANGE_EVENT:
	// 	return "ViewChangeEvent"
	// case replication.XA_PREPARE_LOG_EVENT:
	// 	return "XAPrepareLogEvent"

	default:
		return
	}
}

func handleQueryEvent(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.QueryEvent)
	fmt.Println(ev)
	if string(ev.Schema) == "" {
		apputil.PrettyPrint(e)
		fmt.Println("schema begin")
		apputil.PrettyPrint(string(ev.Schema))
		fmt.Println("schema end")
		e.Dump(os.Stdout)
		return
	}
	FlushDBTables(string(ev.Schema))
	apputil.PrettyPrint(DBTables)
}
func handleUpdateEventV1(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.RowsEvent)
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)
	ok := FilterTable(dbName, tableName)
	if !ok {
		return
	}
	var streams []models.DddEventStream

	stream := &models.DddEventStream{}
	stream.DbName = dbName
	stream.TableName = tableName
	stream.TransactionTag = ""
	stream.EventType = 2 // 此处是canal定义，和原mysql binlog event type 不同

	for i := 0; i < len(ev.Rows); i = i + 2 {
		var allColumns []string
		var updatedColumns []string
		updatedData := make(map[string]interface{})

		next := i + 1
		tableSchema := DBTables[string(ev.Table.Schema)+"."+string(ev.Table.Table)]
		for idx, value := range ev.Rows[i] {
			// fmt.Println("i = ", idx)
			// fmt.Println("name = ", tableSchema[idx])
			// fmt.Println("value = ", value)
			// fmt.Printf("%+v", ev.Rows[next][idx])
			// fmt.Println()
			// fmt.Printf("%+T", ev.Rows[next][idx])
			// fmt.Println()
			allColumns = append(allColumns, tableSchema[idx])

			// go类型断言
			// https://www.jianshu.com/p/787cf3a41ccb
			// mysql 反回字段interface类型， 获取value参考
			// /Users/owen/go/pkg/mod/github.com/go-xorm/xorm@v0.7.9/session_query.go
			if !cmp.Equal(value, ev.Rows[next][idx]) {
				updatedColumns = append(updatedColumns, tableSchema[idx])
				//strValue := fmt.Sprintf("%s", ev.Rows[next][idx])
				strValue := getValueString(ev.Rows[next][idx])
				updatedData[tableSchema[idx]] = strValue
			}
		}

		stream.Columns = strings.Join(allColumns, ",")
		stream.UpdateColumns = strings.Join(updatedColumns, ",")

		b, _ := json.Marshal(updatedData)
		stream.UpdateValue = string(b)

		streams = append(streams, *stream)
	}
	gorm := conn.GetEventGorm()
	gorm.Table("ddd_event_stream").Create(&streams)
}

func handleWriteRowsEventV1(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.RowsEvent)
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)
	ok := FilterTable(dbName, tableName)
	if !ok {
		return
	}

	var streams []models.DddEventStream

	stream := &models.DddEventStream{}
	stream.DbName = dbName
	stream.TableName = tableName
	stream.TransactionTag = ""
	stream.EventType = 1 // 此处是canal定义，和原mysql binlog event type 不同

	for i := 0; i < len(ev.Rows); i++ {
		var allColumns []string
		var updatedColumns []string
		updatedData := make(map[string]interface{})

		tableSchema := DBTables[string(ev.Table.Schema)+"."+string(ev.Table.Table)]
		for idx, value := range ev.Rows[i] {
			allColumns = append(allColumns, tableSchema[idx])
			updatedColumns = append(updatedColumns, tableSchema[idx])
			updatedData[tableSchema[idx]] = value
		}

		stream.Columns = strings.Join(allColumns, ",")
		stream.UpdateColumns = strings.Join(updatedColumns, ",")

		b, _ := json.Marshal(updatedData)
		stream.UpdateValue = string(b)

		streams = append(streams, *stream)
	}
	gorm := conn.GetEventGorm()
	gorm.Table("ddd_event_stream").Create(&streams)
}

func handleDeleteRowsEventV1(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.RowsEvent)
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)
	ok := FilterTable(dbName, tableName)
	if !ok {
		return
	}

	var streams []models.DddEventStream

	stream := &models.DddEventStream{}
	stream.DbName = dbName
	stream.TableName = tableName
	stream.TransactionTag = ""
	stream.EventType = 3 // 此处是canal定义，和原mysql binlog event type 不同

	for i := 0; i < len(ev.Rows); i++ {
		var allColumns []string
		var updatedColumns []string
		updatedData := make(map[string]interface{})

		tableSchema := DBTables[string(ev.Table.Schema)+"."+string(ev.Table.Table)]
		for idx, value := range ev.Rows[i] {
			allColumns = append(allColumns, tableSchema[idx])
			updatedColumns = append(updatedColumns, tableSchema[idx])
			updatedData[tableSchema[idx]] = value
		}

		stream.Columns = strings.Join(allColumns, ",")
		stream.UpdateColumns = strings.Join(updatedColumns, ",")

		b, _ := json.Marshal(updatedData)
		stream.UpdateValue = string(b)

		streams = append(streams, *stream)
	}
	gorm := conn.GetEventGorm()
	gorm.Table("ddd_event_stream").Create(&streams)
}

func getValueString(v interface{}) string {
	rawValue := reflect.Indirect(reflect.ValueOf(v))

	stringV, _ := value2String(&rawValue)
	return stringV
}

// 如果早知这样，直接用xorm就好？
// 另一种类型断言
// https://www.jianshu.com/p/787cf3a41ccb
func value2String(rawValue *reflect.Value) (str string, err error) {
	aa := reflect.TypeOf((*rawValue).Interface())
	vv := reflect.ValueOf((*rawValue).Interface())
	switch aa.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str = strconv.FormatInt(vv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str = strconv.FormatUint(vv.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		str = strconv.FormatFloat(vv.Float(), 'f', -1, 64)
	case reflect.String:
		str = vv.String()
	case reflect.Array, reflect.Slice:
		switch aa.Elem().Kind() {
		case reflect.Uint8:
			data := rawValue.Interface().([]byte)
			str = string(data)
			if str == "\x00" {
				str = "0"
			}
		default:
			err = fmt.Errorf("Unsupported struct type %v", vv.Type().Name())
		}
	// time type
	case reflect.Struct:
		if aa.ConvertibleTo(core.TimeType) {
			str = vv.Convert(core.TimeType).Interface().(time.Time).Format(time.RFC3339Nano)
		} else {
			err = fmt.Errorf("Unsupported struct type %v", vv.Type().Name())
		}
	case reflect.Bool:
		str = strconv.FormatBool(vv.Bool())
	case reflect.Complex128, reflect.Complex64:
		str = fmt.Sprintf("%v", vv.Complex())
	/* TODO: unsupported types below
	   case reflect.Map:
	   case reflect.Ptr:
	   case reflect.Uintptr:
	   case reflect.UnsafePointer:
	   case reflect.Chan, reflect.Func, reflect.Interface:
	*/
	default:
		err = fmt.Errorf("Unsupported struct type %v", vv.Type().Name())
	}
	return
}
