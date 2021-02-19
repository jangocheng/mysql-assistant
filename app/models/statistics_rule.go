package models

type StatisticsRule struct {
	StatisticsRuleId int       `json:"statistics_rule_id"`
	DbName           string    `json:"db_name"`
	TableName        string    `json:"table_name"`
	FieldName        string    `json:"field_name"`
	IsDeleted        int       `json:"is_deleted"`
	CreatedAt        DateTime `json:"created_at"`
	UpdatedAt        DateTime `json:"updated_at"`
	DeletedAt        DateTime `json:"deleted_at"`
}
