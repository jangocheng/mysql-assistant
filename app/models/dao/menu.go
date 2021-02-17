package dao

import (
	"owen2020/app/models"
	"owen2020/conn"
)

func GetMenuList() []models.Menu {
	gorm := conn.GetEventGorm()

	results := []models.Menu{}

	err := gorm.Debug().Table("menu").Where("is_deleted= ?", 0).Find(&results).Error
	if err != nil {
		return nil
	}

	return results
}
