package models

type StateAbnormal struct {
	StateAbnormalId int       `xorm:"not null pk autoincr INT(11)"`
	AbnormalType    int       `xorm:"default 0 INT(11)"`
	EventType       int       `xorm:"default 0 INT(11)"`
	DbName          string    `xorm:"default '' VARCHAR(255)"`
	TableName       string    `xorm:"default '' VARCHAR(255)"`
	FieldName       string    `xorm:"default '' VARCHAR(255)"`
	StateFrom       string    `xorm:"VARCHAR(255)"`
	StateTo         string    `xorm:"VARCHAR(255)"`
	IsDeleted       int       `xorm:"default 0 TINYINT(4)"`
	CreatedAt       DateTime `gorm:"autoUpdateTime;column:created_at;default:null" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt       DateTime `gorm:"autoUpdateTime;column:updated_at;default:null" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	DeletedAt       DateTime `xorm:"DATETIME"`
}
