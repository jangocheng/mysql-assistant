package models

type StatisticsRule struct {
	StatisticsRuleId int       `xorm:"not null pk autoincr INT(11)"`
	DbName           string    `xorm:"default '' VARCHAR(255)"`
	TableName        string    `xorm:"default '' VARCHAR(255)"`
	FieldName        string    `xorm:"default '' VARCHAR(255)"`
	IsDeleted        int       `xorm:"default 0 TINYINT(4)"`
	CreatedAt        DateTime `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt        DateTime `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	DeletedAt        DateTime `xorm:"DATETIME"`
}
