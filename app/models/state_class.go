package models

type StateClass struct {
	DbName        string    `json:"db_name" xorm:"default '' VARCHAR(255)"`
	TableName     string    `json:"table_name" xorm:"default '' VARCHAR(255)"`
	FieldName     string    `json:"field_name" xorm:"default '' VARCHAR(255)"`
	StateClassId  int       `json:"state_class_id" xorm:"not null pk autoincr INT(11)"`
	StateDescribe string    `json:"state_describe" xorm:"default '' VARCHAR(255)"`
	StateName     string    `json:"state_name" xorm:"default '' VARCHAR(255)"`
	CreatedAt     DateTime `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt     DateTime `json:"updated_at" xorm:"DATETIME"`
	IsDeleted     int       `json:"is_deleted" xorm:"default 0 TINYINT(4)"`
	DeletedAt     DateTime `json:"deleted_at" xorm:"DATETIME"`
}
