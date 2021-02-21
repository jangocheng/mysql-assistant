package handle_binlog

import (
	"fmt"
	"github.com/siddontang/go-mysql/replication"
	"owen2020/cmd/command/handle_binlog/common"
)

func handleQueryEvent(e *replication.BinlogEvent) {
	ev, _ := e.Event.(*replication.QueryEvent)
	if string(ev.Schema) == "" {
		//apputil.PrettyPrint(e)
		fmt.Println("event schema is empty")
		//apputil.PrettyPrint(string(ev.Schema))
		//fmt.Println("schema end")
		//e.Dump(os.Stdout)
		return
	}
	common.FlushDBTables(string(ev.Schema))
}
