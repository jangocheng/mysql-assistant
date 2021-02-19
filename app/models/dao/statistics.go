package dao

import (
	"owen2020/app/models"
	"owen2020/conn"
)

func GetStatisticsRuleList() ([]models.StatisticsRule, error) {
	list := []models.StatisticsRule{}

	gorm := conn.GetEventGorm()
	err := gorm.Table("statistics_rule").Where("is_deleted = 0").Find(&list).Error

	return list, err
}
