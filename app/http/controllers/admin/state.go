package admin

import (
	"github.com/gin-gonic/gin"
	"owen2020/app/apputil"
	"owen2020/app/models"
	"owen2020/app/reqt"
	"owen2020/app/resp/out"
	"owen2020/conn"
	"strconv"
	"time"
)

func GetStateClassList(c *gin.Context) {
	pageParams := &reqt.PageParam{}
	c.ShouldBindQuery(&pageParams)

	var total int64
	list := []models.StateClass{}

	db := conn.GetEventGorm()
	query := db.Table("state_class").Where("is_deleted = ?", 0)

	if err := query.Count(&total).Error; err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
		return
	}

	err := query.Order("state_class_id desc").Limit(pageParams.Limit()).Offset(pageParams.Offset()).Find(&list).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(gin.H{"total": total, "rows": list}).JSONOK(c)
}

func GetStateClassInfo(c *gin.Context) {
	id := c.Param("id")
	info := models.StateClass{}

	db := conn.GetEventGorm()
	query := db.Table("state_class").Where("is_deleted = 0 and state_class_id = ?", id)

	err := query.First(&info).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(info).JSONOK(c)
}

func GetStateInfo(c *gin.Context) {
	id := c.Param("id")
	info := models.State{}

	db := conn.GetEventGorm()
	query := db.Table("state").Where("is_deleted = 0 and state_id = ?", id)

	err := query.First(&info).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(info).JSONOK(c)
}

//AddStateClass 添加
func AddStateClass(c *gin.Context) {
	info := models.StateClass{}

	err := apputil.ShouldBindOrError(c, &info)
	if err != nil {
		return
	}

	db := conn.GetEventGorm()
	err = db.Table("state_class").Create(&info).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(info).JSONOK(c)
}

//EditStateClass 编辑
func EditStateClass(c *gin.Context) {
	id := c.Param("id")

	info := models.StateClass{}
	err := apputil.ShouldBindOrError(c, &info)
	if err != nil {
		return
	}
	info.StateClassId,_ = strconv.Atoi(id)

	db := conn.GetEventGorm()
	session := db.Table("state_class").Select("*").Where("state_class_id = ?", id).UpdateColumns(info)
	//stmt := session.Statement
	//fmt.Println("sql is :", stmt.SQL.String())
	err = session.Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(info).JSONOK(c)
}

func AddState(c *gin.Context) {
	info := models.State{}
	err := apputil.ShouldBindOrError(c, &info)

	db := conn.GetEventGorm()
	err = db.Table("state").Create(&info).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(info).JSONOK(c)
}

func EditState(c *gin.Context) {
	id := c.Param("id")

	info := models.State{}
	err := apputil.ShouldBindOrError(c, &info)
	if err != nil {
		return
	}
	info.StateId ,_ = strconv.Atoi(id)

	db := conn.GetEventGorm()
	session := db.Table("state").Select("state_class_id,state_value,state_value_desc").Where("state_id = ?", id).UpdateColumns(info)
	err = session.Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(info).JSONOK(c)
}

func GetStateList(c *gin.Context) {
	pageParams := &reqt.PageParam{}
	c.ShouldBindQuery(&pageParams)

	var stateClassId int
	queryId := c.Query("state_class_id")
	if queryId != "" {
		stateClassId, _ = strconv.Atoi(queryId)
	}

	var total int64
	list := []models.State{}

	db := conn.GetEventGorm()
	query := db.Table("state").Where("is_deleted = ?", 0)
	if stateClassId != 0 {
		query.Where("state_class_id", stateClassId)
	}

	if err := query.Count(&total).Error; err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
		return
	}

	err := query.Order("state_id desc").Limit(pageParams.Limit()).Offset(pageParams.Offset()).Find(&list).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(gin.H{"total": total, "rows": list}).JSONOK(c)
}

func GetStateDirectionList(c *gin.Context) {
	pageParams := &reqt.PageParam{}
	c.ShouldBindQuery(&pageParams)

	var stateClassId int
	queryId := c.Query("state_class_id")
	if queryId != "" {
		stateClassId, _ = strconv.Atoi(queryId)
	}

	var total int64
	list := []models.StateDirection{}

	db := conn.GetEventGorm()
	query := db.Table("state_direction").Where("is_deleted = ?", 0)
	if stateClassId != 0 {
		query.Where("state_class_id", stateClassId)
	}

	if err := query.Count(&total).Error; err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
		return
	}

	err := query.Order("state_direction_id desc").Limit(pageParams.Limit()).Offset(pageParams.Offset()).Find(&list).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(gin.H{"total": total, "rows": list}).JSONOK(c)
}

func GetStateAbnormalList(c *gin.Context) {
	pageParams := &reqt.PageParam{}
	c.ShouldBindQuery(&pageParams)

	var total int64
	list := []models.StateAbnormal{}

	db := conn.GetEventGorm()
	query := db.Table("state_abnormal").Where("is_deleted = ?", 0)

	if err := query.Count(&total).Error; err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
		return
	}

	err := query.Order("state_abnormal_id desc").Limit(pageParams.Limit()).Offset(pageParams.Offset()).Find(&list).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(gin.H{"total": total, "rows": list}).JSONOK(c)
}

func DeleteStatClass(c *gin.Context) {
	id := c.Param("id")

	gorm := conn.GetEventGorm()
	query := gorm.Table("state_class").Where("state_class_id = ?", id)

	err := query.Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess("").JSONOK(c)
}

func DeleteStatAbnormal(c *gin.Context) {
	id := c.Param("id")

	gorm := conn.GetEventGorm()
	query := gorm.Table("state_abnormal").Where("state_abnormal_id = ?", id)

	err := query.Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess("").JSONOK(c)
}

func DeleteStat(c *gin.Context) {
	id := c.Param("id")

	gorm := conn.GetEventGorm()
	query := gorm.Table("state").Where("state_id = ?", id)

	err := query.Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess("").JSONOK(c)
}

func DeleteStatDirection(c *gin.Context) {
	id := c.Param("id")

	gorm := conn.GetEventGorm()
	query := gorm.Table("state_direction").Where("state_direction_id = ?", id)

	err := query.Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess("").JSONOK(c)
}
