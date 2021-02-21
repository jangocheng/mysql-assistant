package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"owen2020/app/apputil/foundations"
	"path/filepath"

	"owen2020/cmd/command"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

// 文档： https://github.com/urfave/cli/blob/master/docs/v2/manual.md#arguments
// Get the parent path https://stackoverflow.com/questions/48570228/get-the-parent-path
// Flags 配置是全局的， 如指定.env配置，   指定单独的配置项等
func main() {
	dir, _ := os.Getwd()
	rootDir, err := getRootDir(dir)
	fmt.Println(rootDir)
	err = godotenv.Load(rootDir + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &cli.App{
		Commands: []*cli.Command{
			//{
			//	Name:    "dev",
			//	Aliases: []string{"c"},
			//	Usage:   "开发运行调试",
			//	Action:  command.Dev,
			//},
			{
				Name:   "binlog-start",
				Usage:  "开始消费mysql binlog",
				Action: command.StartBinlogClient,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "host", Value: "127.0.0.1"},
					&cli.IntFlag{Name: "port", Value: 3306},
					&cli.StringFlag{Name: "username", Value: "root"},
					&cli.StringFlag{Name: "password", Value: ""},
					&cli.StringFlag{Name: "filter", Value: ".*\\..*"},
					&cli.IntFlag{Name: "server_id", Value: 100},
				},
			},
			{
				Name:   "web-start",
				Usage:  "启动web服务器",
				Action: command.StartWebServer,
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getRootDir(dir string) (string, error) {
	if dir == "" {
		return "", errors.New("未能定位根目录")
	}

	if foundations.CheckFileIsExist(dir + "/.env") {
		return dir, nil
	}

	return getRootDir(filepath.Dir(dir))
}
