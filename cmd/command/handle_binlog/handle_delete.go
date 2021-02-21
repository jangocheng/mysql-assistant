package handle_binlog

import (
	"encoding/json"
	"github.com/siddontang/go-mysql/replication"
	"os"
	"owen2020/app/models"
	"owen2020/cmd/command/handle_binlog/common"
	"owen2020/cmd/command/handle_binlog/store"
	"strings"
)

func handleDeleteRowsEventV1(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.RowsEvent)

	if os.Getenv("ENABLE_DATA_STATISTICS") == "yes" {
		go deleteRoutineStatistics(ev)
	}

	if os.Getenv("ENABLE_MODEL_STREAM") == "yes" {
		insertRoutineModelStream(ev)
	}
}

func deleteRoutineStatistics(ev *replication.RowsEvent) {
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)

	store.StatisticsIncrease(dbName, tableName, "", store.INSERT, 1)
}

func deleteRoutineModelStream(ev *replication.RowsEvent) {
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)
	
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

		tableSchema := common.DBTables[string(ev.Table.Schema)+"."+string(ev.Table.Table)]
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

	store.StreamAddRows(streams)
}
