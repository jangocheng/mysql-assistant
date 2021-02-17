package dao

import (
	"owen2020/app/apputil"
	"owen2020/app/models"
	"owen2020/conn"
)

func GetStateClassList() ([]models.StateClass, error) {
	stateClassList := []models.StateClass{}
	gorm := conn.GetEventGorm()
	apputil.PrettyPrint(gorm)
	err := gorm.Table("state_class").Where("status = ? and is_deleted = 0", 1).Find(&stateClassList).Error

	return stateClassList, err
}

