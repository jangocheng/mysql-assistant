package handle_binlog

import (
	"encoding/json"
	"fmt"
	"github.com/siddontang/go-mysql/replication"
	"os"
	"owen2020/app/models"
	"owen2020/conn"
	"strings"
)

func handleDeleteRowsEventV1(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.RowsEvent)
	dbName := string(ev.Table.Schema)
	tableName := string(ev.Table.Table)
	ok := FilterTable(dbName, tableName)
	if !ok {
		fmt.Println("skip delete", dbName, ".", tableName)
		return
	}

	if os.Getenv("ENABLE_DATA_STATISTICS") == "yes" {
		StatisticsIncrease(dbName, tableName, "", DELETE, 1)
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
