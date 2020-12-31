package models

type DddEventStream struct {
	DddEventStreamId  int      `xorm:"not null pk autoincr INT(11)" json:"ddd_event_stream_id"`
	DbName            string   `xorm:"VARCHAR(255)" json:"db_name"`
	TableName         string   `xorm:"VARCHAR(255)" json:"table_name"`
	TransactionTag    string   `xorm:"VARCHAR(64)" json:"transaction_tag"`
	EventType         int      `xorm:"default -100 INT(11)" json:"event_type"`
	Columns           string   `xorm:"TEXT" json:"columns"`
	UpdateColumns     string   `xorm:"comment('更新的字段') TEXT" json:"update_columns"`
	UpdateValue       string   `xorm:"comment('更新字段的值') TEXT" json:"update_value"`
	IgnoreColumnValue string   `xorm:"comment('忽略的字段值') TEXT" json:"ignore_column_value"`
	Comment           string   `xorm:"TEXT" json:"comment"`
	IsDeleted         int      `xorm:"default 0 TINYINT(4)" json:"is_deleted"`
	DeletedAt         DateTime `xorm:"DATETIME" json:"deleted_at"`
	CreatedAt         DateTime `gorm:"autoUpdateTime;column:created_at;default:null" xorm:"default 'CURRENT_TIMESTAMP' DATETIME" json:"created_at"`
	UpdatedAt         DateTime `gorm:"autoUpdateTime;column:updated_at;default:null" xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME" json:"updated_at"`
}
