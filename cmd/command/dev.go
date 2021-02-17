package command

import (
	"fmt"
	"log"
	"os/exec"
	"owen2020/app/apputil"
	"owen2020/app/models/dao"
	"owen2020/conn"
	"runtime"

	"github.com/urfave/cli/v2"
)

//How to pretty print a Golang structure? [duplicate]
// https://stackoverflow.com/questions/56242013/how-to-pretty-print-a-golang-structure

// [Golang] Pretty Print Variable (struct, map, array, slice)
// https://siongui.github.io/2016/01/30/go-pretty-print-variable/

func Dev(c *cli.Context) error {
	// 初始化event数据库链接池
	conn.InitEventGormPool()

	// gorm := conn.GetDefaultGorm()
	// mysqlutil.GetMysqlPosition()
	// apputil.PrettyPrint(dbs)

	// for _, db := range dbs {
	// 	tables := handle_binlog.GetTableNames(gorm, db)
	// 	apputil.PrettyPrint(tables)
	// }

	// apputil.PrettyPrint(config.GetServers())
	// binlog.FlushTableIdentifierNameMap("codeper", "user")
	// binlog.InitDBTables("codeper")

	// apputil.PrettyPrint(binlog.DBTables)
	// ok, err := regexp.Match(filter, []byte("codeper"+"."+"user"))
	// fmt.Println(ok)
	// fmt.Println(err)
	// tables := binlog.GetTableNames("codeper")

	list, _ := dao.GetStateClassList()
	apputil.PrettyPrint(list)
	apputil.PrettyPrint("aaa")
	return nil
}

func OpenBrowser(c *cli.Context) error {

	url := "http://www.baidu.com"

	openbrowser(url)

	return nil
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
