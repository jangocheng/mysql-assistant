package activity

import (
	"fmt"
	"owen2020/app/apputil"
	"owen2020/app/models"
	"owen2020/app/resp/out"
	"owen2020/conn"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "gorm.io/gorm"
)

func EnterAdd(c *gin.Context) {
	activityEnter := &models.ActivityEnter{}
	err := apputil.ShouldBindOrError(c, &activityEnter)
	if nil != err {
		return
	}

	fmt.Printf("%+v", activityEnter)
	memberID, has := c.Get("member_id")
	if has {
		activityEnter.MemberId, _ = strconv.Atoi(memberID.(string))
	}

	gorm := conn.GetDefaultGorm()
	result := gorm.Table("activity_enter").Create(&activityEnter)
	if result.Error != nil {
		out.NewError(701, result.Error.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(activityEnter).JSONOK(c)
	return
}

func EnterList(c *gin.Context) {
	memberID, has := c.Get("member_id")
	if !has {
		out.NewError(701, "用户未登录").JSONOK(c)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	var enterList []models.ActivityEnter

	gorm := conn.GetGormWithLog()
	session := gorm.Table("activity_enter").Where("member_id = ? ", memberID)

	var total int64

	if err := session.Count(&total).Error; err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
		return
	}
	err := session.Limit(pageSize).Offset((page - 1) * pageSize).Order("activity_enter_id desc").Find(&enterList).Error
	if nil != err {
		out.NewError(600, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(gin.H{"total": total, "rows": enterList}).JSONOK(c)
}

//EnterListByActivity 公开的报名人列表
func EnterListByActivity(c *gin.Context) {
	// memberId, has := c.Get("member_id")
	// if !has {
	// 	out.NewError(701, "用户未登录").JSONOK(c)
	// 	return
	// }

	activityId := c.Query("activity_id")
	var enterList []models.ActivityEnter

	gorm := conn.GetGormWithLog()
	session := gorm.Table("activity_enter").Where("activity_id =? and status>=0", activityId)

	err := session.Order("activity_enter_id desc").Find(&enterList).Error
	if nil != err {
		out.NewError(600, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(gin.H{"rows": enterList}).JSONOK(c)
}

func getEnterMemberIds(list []models.ActivityEnter) []int {
	s := make([]int, 0, len(list))
	for _, v := range list {
		s = append(s, v.MemberId)
	}
	return s
}

func getEnterActivityIds(list []models.ActivityEnter) []int {
	s := make([]int, 0, len(list))
	for _, v := range list {
		s = append(s, v.ActivityId)
	}
	return s
}

func EnterInfo(c *gin.Context) {
	id := c.Param("id")
	// idInt, _ := strconv.Atoi(id)
	memberID, _ := c.Get("member_id")
	// memberIDStr, _ := strconv.Atoi(memberID.(string))
	var enterInfo models.ActivityEnter

	// gorm := conn.GetGormWithConfig(&gorm.Config{DryRun: true})
	// stmt := gorm.Table("activity_enter").Where("member_id=? and activity_enter_id=?", memberID, id).First(&enterInfo).Statement
	// fmt.Println(stmt.SQL.String())
	// fmt.Println(stmt.Vars)

	gorm := conn.GetDefaultGorm()
	err := gorm.Table("activity_enter").Where("member_id=? and activity_enter_id=?", memberID, id).First(&enterInfo).Error
	if nil != err {
		out.NewError(600, err.Error()).JSONOK(c)
		return
	}

	// respInfo := resputil.GenInfoByActInfo(enterInfo)
	out.NewSuccess(enterInfo).JSONOK(c)

}

func EnterMemberList(c *gin.Context) {

}

func EnterMemberInfo(c *gin.Context) {

}
