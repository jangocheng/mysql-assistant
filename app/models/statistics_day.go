package models

type StatisticsDay struct {
	StatisticsDayId  int      `gorm:"primaryKey" json:"statistics_day_id"`
	StatisticsRuleId int      `json:"statistics_rule_id"`
	StatisticsDay    Date     `json:"statistics_day"`
	DbName           string   `json:"db_name"`
	TableName        string   `json:"table_name"`
	FieldName        string   `json:"field_name"`
	InsertTimes      int      `json:"insert_times"`
	UpdateTimes      int      `json:"update_times"`
	DeleteTimes      int      `json:"delete_times"`
	IsDeleted        int      `json:"is_deleted"`
	CreatedAt        DateTime `gorm:"autoUpdateTime;column:created_at;default:null" json:"created_at"`
	UpdatedAt        DateTime `gorm:"autoUpdateTime;column:updated_at;default:null" json:"updated_at"`
	DeletedAt        DateTime `json:"deleted_at"`
}
