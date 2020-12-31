package command

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"owen2020/app/models"
	"owen2020/conn"
	"strings"
	"time"

	"runtime/debug"

	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli/v2"

	uuid "github.com/satori/go.uuid"
	"github.com/withlin/canal-go/client"
	protocol "github.com/withlin/canal-go/protocol"
)

/*
run.sh -e canal.instance.master.address=host:3306 \
         -e canal.instance.dbUsername= \
         -e canal.instance.dbPassword= \
         -e canal.instance.connectionCharset=UTF-8 \
         -e canal.instance.tsdb.enable=true \
         -e canal.instance.gtidon=false \
         -e canal.destinations=test \
         -e canal.instance.filter.regex=".codeper\\..*"
*/

var uniqid string

func CanalClient(c *cli.Context) error {
	address := "127.0.0.1"
	port := 11111
	username := "canle"
	password := "owen@libu2018"
	destination := "test"
	// destination := "example"
	soTimeOut := int32(0)
	idleTimeOut := int32(0)
	connector := client.NewSimpleCanalConnector(address, port, username, password, destination, soTimeOut, idleTimeOut)

	err := connector.Connect()
	if err != nil {
		log.Println(err)
		// fmt.Printf("%s", debug.Stack())
		debug.PrintStack()
		os.Exit(1)
	}

	// err = connector.Subscribe(".*\\..*")
	err = connector.Subscribe("codeper\\..*")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	gorm := conn.GetEventGorm()
	for {
		message, err := connector.Get(100, nil, nil)
		if err != nil {
			fmt.Println("===获取消息误错===")
			log.Println(err)
			os.Exit(1)
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("===没有数据了===")
			continue
		}

		// printEntry(message.Entries)
		for _, entity := range message.Entries {
			info := extractElementInfo(entity)
			if info != nil {
				gorm.Table("ddd_event_stream").Create(info)
			}
		}

	}

	return nil
}

func extractElementInfo(entry protocol.Entry) []models.DddEventStream {

	switch entry.GetEntryType() {
	case protocol.EntryType_TRANSACTIONBEGIN: // 事务开启
		uniqid = uuid.NewV4().String()
		// 设置事务标识
		return nil
	case protocol.EntryType_TRANSACTIONEND: // 事件结束
		uniqid = ""
		return nil
	}
	rowChange := new(protocol.RowChange)
	err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	var streams []models.DddEventStream

	eventType := rowChange.GetEventType()

	for _, rowData := range rowChange.GetRowDatas() {
		stream := &models.DddEventStream{}
		stream.DbName = entry.Header.SchemaName
		stream.TableName = entry.Header.TableName
		stream.TransactionTag = uniqid
		stream.EventType = int(eventType)

		var allColumns []string
		var updatedColumns []string
		updatedData := make(map[string]interface{})

		afterColumns := rowData.GetAfterColumns()

		switch eventType {
		case protocol.EventType_INSERT:
			for _, proColumn := range afterColumns {
				allColumns = append(allColumns, proColumn.GetName())
				updatedColumns = append(updatedColumns, proColumn.GetName())
				updatedData[proColumn.GetName()] = proColumn.GetValue()
			}

		case protocol.EventType_DELETE:
			for _, proColumn := range afterColumns {
				allColumns = append(allColumns, proColumn.GetName())
				updatedColumns = append(updatedColumns, proColumn.GetName())
				updatedData[proColumn.GetName()] = proColumn.GetValue()
			}

		case protocol.EventType_UPDATE:
			for _, proColumn := range afterColumns {
				allColumns = append(allColumns, proColumn.GetName())

				if proColumn.Updated {
					updatedColumns = append(updatedColumns, proColumn.GetName())
					updatedData[proColumn.GetName()] = proColumn.GetValue()
				}
			}

		}

		stream.Columns = strings.Join(allColumns, ",")
		stream.UpdateColumns = strings.Join(updatedColumns, ",")

		b, _ := json.Marshal(updatedData)
		stream.UpdateValue = string(b)

		streams = append(streams, *stream)
	}
	return streams
}
