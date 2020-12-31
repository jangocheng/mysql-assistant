package dao

import (
	"owen2020/app/models"
	"owen2020/conn"
)

func GetMenuList() []models.Menu {
	gorm := conn.GetEventGorm()

	results := []models.Menu{}

	err := gorm.Table("menu").Where("status > ? and is_deleted= ?", 0, 0).Find(&results).Error
	if err != nil {
		return nil
	}

	return results
}
