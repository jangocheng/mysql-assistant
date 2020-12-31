package log

import (
	"github.com/gin-gonic/gin"
	// log "github.com/sirupsen/logrus"
	"owen2020/app/apputil/applog"
)

func TestLog(c *gin.Context) {
	// fileLog := log.New().SetOutput()
	// log.WithFields(log.Fields{"哈哈": "黑黑"}).Info("A group of walrus emerges from the ocean")
	applog.Logger.Info("记一个中文日志")
}
