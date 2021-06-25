package app

import (
	"os"
	"owen2020/app/apputil/foundations"
	"owen2020/conn"
	"owen2020/routes"
)

func StartWebServer() {
	defer func() {
		routes.RmPidFile()
	}()

	go func() {
		if os.Getenv("APP_ENV") == "local" {
			foundations.OpenBrowser("http://127.0.0.1" + os.Getenv("APP_PORT") + "/admin/entrance/login.html")
		}
	}()

	_ = conn.InitEventGormPool()
	routes.StartGin()
}
