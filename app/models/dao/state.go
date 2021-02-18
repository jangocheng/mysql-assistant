package dao

import (
	"owen2020/app/models"
	"owen2020/conn"
)

func GetStateClassList() ([]models.StateClass, error) {
	stateClassList := []models.StateClass{}
	gorm := conn.GetEventGorm()
	err := gorm.Table("state_class").Where("status = ? and is_deleted = 0", 1).Find(&stateClassList).Error

	return stateClassList, err
}

func GetStateDirectionList(stateClassId int) ([]models.StateDirection, error) {
	directionList := []models.StateDirection{}
	gorm := conn.GetEventGorm()

	err := gorm.Table("state_direction").Where("state_class_id = ? and is_deleted = 0", stateClassId).Find(&directionList).Error

	return directionList, err
}

func GetAllStateDirectionList() ([]models.StateDirection, error) {
	directionList := []models.StateDirection{}
	gorm := conn.GetEventGorm()

	err := gorm.Table("state_direction").Where("is_deleted = 0").Order("state_class_id").Find(&directionList).Error

	return directionList, err
}
