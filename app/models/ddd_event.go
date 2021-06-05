package models

type DddEvent struct {
	DddEventId   int      `gorm:"primaryKey" xorm:"not null pk autoincr INT(11)" json:"ddd_event_id" form:"ddd_event_id"`
	EventType    string   `xorm:"default 'mysql-update' VARCHAR(64)" json:"event_type" form:"event_type"`
	EventTag     string   `xorm:"comment('事件标签') VARCHAR(255)" json:"event_tag" form:"event_tag"`
	EventName    string   `xorm:"comment('事件名称') VARCHAR(255)" json:"event_name" form:"event_name" binding:"required" validate:"required,min=2"`
	StreamIds    string   `xorm:"comment('事件包含的mysql操作') VARCHAR(255)" json:"stream_ids" form:"stream_ids"`
	EventVersion string   `xorm:"VARCHAR(255)" json:"event_version" form:"event_version"`
	EventLink    string   `xorm:"VARCHAR(255)" json:"event_link" form:"event_link"`
	Comment      string   `xorm:"comment('事件说明') TEXT" json:"comment" form:"comment"`
	IsDeleted    int      `xorm:"default 0 INT(4)" json:"is_deleted" form:"is_deleted"`
	DeletedAt    DateTime `gorm:"autoCreateTime" xorm:"DATETIME" json:"deleted_at" form:"deleted_at"`
	CreatedAt    DateTime `gorm:"autoUpdateTime;column:created_at;default:null" xorm:"default 'CURRENT_TIMESTAMP' DATETIME" json:"created_at" form:"created_at"`
	UpdatedAt    DateTime `gorm:"autoUpdateTime;column:updated_at;default:null" xorm:"default 'CURRENT_TIMESTAMP' DATETIME" json:"updated_at" form:"updated_at"`
}
