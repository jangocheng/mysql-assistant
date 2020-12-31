package dao

import (
	"owen2020/app/models"
	"owen2020/conn"
	"strings"
)

func GetStreamListByIds(ids []string) ([]models.DddEventStream, error) {
	streamList := []models.DddEventStream{}
	gorm := conn.GetEventGorm()
	err := gorm.Table("ddd_event_stream").Where("ddd_event_stream_id in ?", ids).Group("db_name").Group("table_name").Find(&streamList).Error

	return streamList, err
}

func GetStreamListByEventId(eventId string) ([]models.DddEventStream, error) {
	streamList := []models.DddEventStream{}

	eventInfo := &models.DddEvent{}
	gorm := conn.GetEventGorm()
	err := gorm.Table("ddd_event").Where("ddd_event_id = ?", eventId).First(&eventInfo).Error
	if err != nil {
		return nil, err
	}

	ids := strings.Split(eventInfo.StreamIds, ",")
	err = gorm.Table("ddd_event_stream").Where("ddd_event_stream_id in ?", ids).Group("db_name").Group("table_name").Find(&streamList).Error
	if err != nil {
		return nil, err
	}

	return streamList, nil
}

func GetStreamListByTransactionId(transactionId string) ([]models.DddEventStream, error) {
	streamList := []models.DddEventStream{}
	gorm := conn.GetEventGorm()
	err := gorm.Table("ddd_event_stream").Where("transaction_tag = ?", transactionId).Group("db_name").Group("table_name").Find(&streamList).Error

	return streamList, err
}
