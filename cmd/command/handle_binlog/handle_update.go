package handle_binlog

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/siddontang/go-mysql/replication"
	"os"
	"owen2020/app/models"
	"owen2020/conn"
	"strings"
)

func handleUpdateEventV1(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.RowsEvent)

	if os.Getenv("ENABLE_DATA_STATISTICS") == "yes" {
		go updateRoutineStatistics(ev)
	}

	if os.Getenv("ENABLE_CHECK_STATE") == "yes" {
		go updateRoutineStatRule(ev)
	}

	if os.Getenv("ENABLE_MODEL_STREAM") == "yes" {
		updateRoutineModelStream(ev)
	}
}

func updateRoutineModelStream(ev *replication.RowsEvent) {
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)
	ok := FilterTable(dbName, tableName)
	if !ok {
		fmt.Println("skip update", dbName, ".", tableName)
		return
	}
	tableSchema := DBTables[dbName+"."+tableName]

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
		for idx, value := range ev.Rows[i] {
			fieldName := tableSchema[idx]
			allColumns = append(allColumns, fieldName)

			// go类型断言
			// https://www.jianshu.com/p/787cf3a41ccb
			// mysql 反回字段interface类型， 获取value参考
			// /Users/owen/go/pkg/mod/github.com/go-xorm/xorm@v0.7.9/session_query.go
			if !cmp.Equal(value, ev.Rows[next][idx]) {
				updatedColumns = append(updatedColumns, fieldName)
				//strValue := fmt.Sprintf("%s", ev.Rows[next][idx])
				strValue := getValueString(ev.Rows[next][idx])
				updatedData[fieldName] = strValue
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

func updateRoutineStatRule(ev *replication.RowsEvent) {
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)
	tableSchema := DBTables[dbName+"."+tableName]

	for i := 0; i < len(ev.Rows); i = i + 2 {
		next := i + 1
		for idx, value := range ev.Rows[i] {
			if cmp.Equal(value, ev.Rows[next][idx]) {
				continue
			}

			fieldName := tableSchema[idx]
			// 校验状态流
			classId, err := GetStatClassId(dbName, tableName, fieldName)
			if nil != err {
				fmt.Println(dbName, tableName, fieldName, err)
				continue
			}

			from, _ := getStringValue(value)
			to, _ := getStringValue(ev.Rows[next][idx])
			check, err := CheckStatDirection(classId, from, to)
			// 流程变更不合规， 做一些通知URL, 钉钉，记录库等
			if !check {
				fmt.Println(dbName, tableName, fieldName, "classId:", classId, "from:", from, "to:", to, err)
				saveStateAbnormal(dbName, tableName, fieldName, from, to)
			}
		}
	}
}
func updateRoutineStatistics(ev *replication.RowsEvent) {
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)
	tableSchema := DBTables[dbName+"."+tableName]

	StatisticsIncrease(dbName, tableName, "", UPDATE, 1)

	for i := 0; i < len(ev.Rows); i = i + 2 {
		next := i + 1
		for idx, value := range ev.Rows[i] {
			if cmp.Equal(value, ev.Rows[next][idx]) {
				continue
			}

			fieldName := tableSchema[idx]
			StatisticsIncrease(dbName, tableName, fieldName, UPDATE, 1)
		}
	}
}

func saveStateAbnormal(dbName string, tableName string, fieldName string, stateFrom string, stateTo string) {
	info := models.StateAbnormal{}
	info.DbName = dbName
	info.TableName = tableName
	info.FieldName = fieldName
	info.StateFrom = stateFrom
	info.StateTo = stateTo

	gorm := conn.GetEventGorm()
	gorm.Table("state_abnormal").Create(&info)
}
