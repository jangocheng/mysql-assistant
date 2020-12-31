package activity

import (
	"fmt"
	"owen2020/app/models"
	"owen2020/app/resp"
	"owen2020/app/resp/out"
	"owen2020/conn"
	"strconv"

	"github.com/gin-gonic/gin"
)

//MyActivityList 我发起组织的活动
func MyActivityList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// id := c.Param("id")
	// idInt, _ := strconv.Atoi(id)

	memberID, _ := c.Get("member_id")
	// memberIDStr, _ := strconv.Atoi(memberID.(string))
	fmt.Println("member_id:", memberID)
	var activityList []models.Activity
	gorm := conn.GetGormWithLog()
	query := gorm.Table("activity").Where("member_id = ?", memberID)
	var total int64

	if err := query.Count(&total).Error; err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
		return
	}
	query.Limit(pageSize).Offset((page - 1) * pageSize).Order("activity_id desc").Find(&activityList)

	respList := resp.GenActivityList(activityList)
	out.NewSuccess(gin.H{"total": total, "rows": respList}).JSONOK(c)
}

//MyActivityInfo 我发起组织的活动详情
func MyActivityInfo(c *gin.Context) {

	id := c.Param("id")
	// idInt, _ := strconv.Atoi(id)
	memberID, _ := c.Get("member_id")
	// memberIDStr, _ := strconv.Atoi(memberID.(string))
	var activityInfo models.Activity

	gorm := conn.GetGormWithLog()
	err := gorm.Table("activity").Where("member_id = ? and activity_id=?", memberID, id).First(&activityInfo).Error

	if nil != err {
		out.NewError(600, err.Error()).JSONOK(c)
		return
	}

	respInfo := resp.GenInfoByActInfo(activityInfo)
	out.NewSuccess(respInfo).JSONOK(c)

}
