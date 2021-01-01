package command

import (
	"github.com/urfave/cli/v2"
	"owen2020/app"
)

func StartWebServer(c *cli.Context) error {
	app.StartWebServer()
	return nil
}
