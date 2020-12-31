package applog

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func init() {
	InitLogger()
}

func InitLogger() {
	pwd, _ := os.Getwd()
	pwd += "/storage/logs/"
	logf, err := rotatelogs.New(
		pwd+"app_%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(pwd+"/app_log"),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}

	Logger = log.New()
	Logger.SetFormatter(&log.JSONFormatter{})
	Logger.SetOutput(logf)
}
