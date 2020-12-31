package admin

import (
	"owen2020/app/models"
	"owen2020/app/reqt"
	"owen2020/app/resp"
	"owen2020/app/resp/out"
	"owen2020/conn"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDddEventStreamList(c *gin.Context) {
	pageParams := &reqt.PageParam{}
	c.ShouldBindQuery(&pageParams)

	eventStreamList := []models.DddEventStream{}
	gorm := conn.GetEventGorm()
	query := gorm.Table("ddd_event_stream").Where("is_deleted = 0")

	var total int64
	if err := query.Count(&total).Error; err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
		return
	}

	err := query.Order("ddd_event_stream_id desc").Limit(pageParams.Limit()).Offset(pageParams.Offset()).Find(&eventStreamList).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	list := []resp.EventStream{}
	for _, v := range eventStreamList {
		list = append(list, resp.GenDiffEntity(v))
	}

	out.NewSuccess(gin.H{"total": total, "rows": list}).JSONOK(c)
}

func DeleteDddEventStream(c *gin.Context) {
	id := c.Param("id")

	gorm := conn.GetEventGorm()
	query := gorm.Table("ddd_event_stream").Where("ddd_event_stream_id = ?", id)

	err := query.Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess("").JSONOK(c)
}
