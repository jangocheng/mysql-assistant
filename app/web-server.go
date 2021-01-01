package app

import (
	"os"
	"owen2020/app/apputil/foundations"
	"owen2020/routes"
)

func StartWebServer() {
	go func() {
		if os.Getenv("APP_ENV") == "local" {
			foundations.OpenBrowser("http://127.0.0.1" + os.Getenv("APP_PORT") + "/admin/entrance/login.html")
		}
	}()

	defer func() {
		routes.RmPidFile()
	}()

	routes.StartGin()
}
