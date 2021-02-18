package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"owen2020/app/apputil"
	"owen2020/cmd/command/handle_binlog"
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
	handle_binlog.InitState()
	apputil.PrettyPrint(handle_binlog.StateClasses)
	apputil.PrettyPrint(handle_binlog.StateClassDirections)

	ok, err := handle_binlog.CheckClassDirection(1, 1, 2)
	apputil.PrettyPrint(ok)
	apputil.PrettyPrint(err)
	ok, err = handle_binlog.CheckClassDirection(2, 1, 2)
	apputil.PrettyPrint(ok)
	fmt.Println(err)
	ok, err = handle_binlog.CheckClassDirection(1, 100, 2)
	apputil.PrettyPrint(ok)
	fmt.Println(err)

	fmt.Println(os.Getenv("ENABLE_CHECK_STATE"))
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
