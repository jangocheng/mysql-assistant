package admin

import (
	"github.com/gin-gonic/gin"
	"owen2020/app/models"
	"owen2020/app/resp/out"
	"owen2020/conn"
)

func GetMenuList(c *gin.Context) {

	list := []models.Menu{}

	var count int64
	db := conn.GetDefaultGorm()
	query := db.Table("menu").Where("status > ? and is_deleted = ?", 0, 0)

	query.Count(&count)
	query.Find(&list)

	out.NewSuccess(gin.H{"total": count, "rows": list}).JSONOK(c)
}
