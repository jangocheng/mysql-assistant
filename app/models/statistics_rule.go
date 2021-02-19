package models

type StatisticsRule struct {
	StatisticsRuleId int       `form:"statistics_rule_id" json:"statistics_rule_id"`
	DbName           string    `form:"db_name" json:"db_name"`
	TableName        string    `form:"table_name" json:"table_name"`
	FieldName        string    `form:"field_name" json:"field_name"`
	IsDeleted        int       `form:"is_deleted" json:"is_deleted"`
	CreatedAt        DateTime `form:"created_at" json:"created_at"`
	UpdatedAt        DateTime `form:"updated_at" json:"updated_at"`
	DeletedAt        DateTime `form:"deleted_at" json:"deleted_at"`
}
