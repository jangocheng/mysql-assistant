package main

import (
	"fmt"
	"log"
	"os"

	"owen2020/cmd/command"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

// 文档： https://github.com/urfave/cli/blob/master/docs/v2/manual.md#arguments
// Get the parent path https://stackoverflow.com/questions/48570228/get-the-parent-path
// Flags 配置是全局的， 如指定.env配置，   指定单独的配置项等
func main() {
	dir, _ := os.Getwd()
	// parent := filepath.Dir(dir)
	fmt.Println(dir)
	err := godotenv.Load(dir + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &cli.App{
		// Flags: []cli.Flag{
		// &cli.StringFlag{
		// 	Name:    "lang",
		// 	Aliases: []string{"l"},
		// 	Value:   "english",
		// 	Usage:   "Language for the greeting",
		// },
		// &cli.StringFlag{
		// 	Name:  "env",
		// 	Value: "../.env",
		// 	// Aliases: []string{"c"},
		// 	Usage:       "specify env file",
		// 	Destination: &envFile,
		// },
		// },
		Commands: []*cli.Command{
			{
				Name:    "dev",
				Aliases: []string{"c"},
				Usage:   "开发运行调试",
				Action:  command.Dev,
			},
			{
				Name: "binlog-start",
				// Aliases: []string{"a"},
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
				Name: "canal-start",
				// Aliases: []string{"a"},
				Usage:  "启动canal客户端",
				Action: command.CanalClient,
			},
			{
				Name: "grpc-server-start",
				// Aliases: []string{"a"},
				Usage:  "启动grpc服务端",
				Action: command.GrpcServerStart,
			},

			{
				Name: "grpc-client-start",
				// Aliases: []string{"a"},
				Usage:  "启动grpc客户端",
				Action: command.GrpcClientStart,
			},
			{
				Name: "zipkin",
				// Aliases: []string{"a"},
				Usage:  "zipkin日志",
				Action: command.ZipkinTest,
			},
			{
				Name: "open-browser",
				// Aliases: []string{"a"},
				Usage:  "打开浏览器",
				Action: command.OpenBrowser,
			},
		},
	}

	//apputil.PrettyPrint(os.Args)
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
