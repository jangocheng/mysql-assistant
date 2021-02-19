package models

type StatisticsDay struct {
	StatisticsDayId  int      `gorm:"primaryKey"`
	StatisticsRuleId int      ``
	StatisticsDay    Date     ``
	DbName           string   ``
	TableName        string   ``
	FieldName        string   ``
	InsertTimes      int      ``
	UpdateTimes      int      ``
	DeleteTimes      int      ``
	IsDeleted        int      ``
	CreatedAt        DateTime `gorm:"autoUpdateTime;column:created_at;default:null"`
	UpdatedAt        DateTime `gorm:"autoUpdateTime;column:updated_at;default:null"`
	DeletedAt        DateTime ``
}
