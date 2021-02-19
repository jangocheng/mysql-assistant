package admin

import (
	"github.com/gin-gonic/gin"
	"owen2020/app/apputil"
	"owen2020/app/models"
	"owen2020/app/models/dao"
	"owen2020/app/reqt"
	"owen2020/app/resp/out"
	"owen2020/conn"
	"time"
)

func GetMenuList(c *gin.Context) {
	pageParams := &reqt.PageParam{}
	c.ShouldBindQuery(&pageParams)

	var total int64
	list := []models.Menu{}

	db := conn.GetEventGorm()
	query := db.Table("menu").Where("is_deleted = ?", 0)

	if err := query.Count(&total).Error; err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
		return
	}

	err := query.Order("menu_id desc").Limit(pageParams.Limit()).Offset(pageParams.Offset()).Find(&list).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(gin.H{"total": total, "rows": list}).JSONOK(c)
}

func GetMenuInfo(c *gin.Context) {
	id := c.Param("id")
	info := models.Menu{}

	db := conn.GetEventGorm()
	query := db.Table("menu").Where("is_deleted = ? and menu_id= ?", 0, id)

	err := query.First(&info).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(info).JSONOK(c)
}

//AddMenu 添加
func AddMenu(c *gin.Context) {
	info := models.Menu{}

	err := apputil.ShouldBindOrError(c, &info)
	if err != nil {
		return
	}

	db := conn.GetEventGorm()
	err = db.Table("menu").Create(&info).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(info).JSONOK(c)
}

//EditMenu 编辑
func EditMenu(c *gin.Context) {
	id := c.Param("id")

	info := models.Menu{}
	err := apputil.ShouldBindOrError(c, &info)
	if err != nil {
		return
	}

	db := conn.GetEventGorm()
	session := db.Table("menu").Where("menu_id = ?", id).UpdateColumns(info)
	err = session.Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(info).JSONOK(c)
}

func DeleteMenu(c *gin.Context) {
	id := c.Param("id")

	gorm := conn.GetEventGorm()
	query := gorm.Table("menu").Where("menu_id = ?", id)

	err := query.Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess("").JSONOK(c)
}

func GetMenuSelectList(c *gin.Context) {
	list := dao.GetMenuList()
	ml := menuList(list)
	//ret := processToTree(list, 0, 0)
	tree := ml.processToTree(0, 0)

	selectData := []menuSelectItem{}

	TreeToSelect(&selectData, tree, 0)

	//apputil.PrettyPrint(tree)
	//apputil.PrettyPrint(selectData)
	out.NewSuccess(selectData).JSONOK(c)
}
