package store

import (
	"owen2020/app/models"
	"owen2020/conn"
)

func SaveStateAbnormal(dbName string, tableName string, fieldName string, stateFrom string, stateTo string) {
	info := models.StateAbnormal{}
	info.DbName = dbName
	info.TableName = tableName
	info.FieldName = fieldName
	info.StateFrom = stateFrom
	info.StateTo = stateTo

	gorm := conn.GetEventGorm()
	gorm.Table("state_abnormal").Create(&info)
}
