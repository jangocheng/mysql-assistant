package dao

import (
	"owen2020/app/models"
	"owen2020/conn"
)

func GetAdminInfo(s string) *models.Administrator {
	db := conn.GetEventGorm()
	info := models.Administrator{}

	err := db.Table("administrator").Where("username = ?", s).First(&info).Error
	if err != nil {
		return nil
	}
	return &info
}
