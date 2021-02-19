package admin

import (
	"github.com/gin-gonic/gin"
	"owen2020/app/models"
	"owen2020/app/reqt"
	"owen2020/app/resp/out"
	"owen2020/conn"
	"time"

)


func GetStatisticRuleList(c *gin.Context) {
	pageParams := &reqt.PageParam{}
	c.ShouldBindQuery(&pageParams)

	var total int64
	list := []models.StatisticsRule{}

	db := conn.GetEventGorm()
	query := db.Table("statistics_rule").Where("is_deleted = ?", 0)

	if err := query.Count(&total).Error; err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
		return
	}

	err := query.Order("statistics_rule_id desc").Limit(pageParams.Limit()).Offset(pageParams.Offset()).Find(&list).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(gin.H{"total": total, "rows": list}).JSONOK(c)
}

func DeleteStatisticRule(c *gin.Context) {
	id := c.Param("id")

	gorm := conn.GetEventGorm()
	query := gorm.Table("statistics_rule").Where("statistics_rule_id = ?", id)

	err := query.Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess("").JSONOK(c)
}

func GetStatisticDayList(c *gin.Context) {
	pageParams := &reqt.PageParam{}
	c.ShouldBindQuery(&pageParams)

	var total int64
	list := []models.StatisticsDay{}

	db := conn.GetEventGorm()
	query := db.Table("statistics_day").Where("is_deleted = ?", 0)

	if err := query.Count(&total).Error; err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
		return
	}

	err := query.Order("statistics_day_id desc").Limit(pageParams.Limit()).Offset(pageParams.Offset()).Find(&list).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(gin.H{"total": total, "rows": list}).JSONOK(c)
}

//func DeleteStatisticDay(c *gin.Context) {
//	id := c.Param("id")
//
//	gorm := conn.GetEventGorm()
//	query := gorm.Table("statistics_day").Where("statistics_day_id = ?", id)
//
//	err := query.Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
//	if err != nil {
//		out.NewError(800, err.Error()).JSONOK(c)
//		return
//	}
//
//	out.NewSuccess("").JSONOK(c)
//}
