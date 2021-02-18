package models

type State struct {
	CreatedAt      DateTime `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	DeletedAt      DateTime `xorm:"DATETIME"`
	IsDeleted      int       `xorm:"default 0 TINYINT(4)"`
	StateClassId   int       `xorm:"INT(11)"`
	StateId        int       `xorm:"not null pk autoincr INT(11)"`
	StateValue     string    `xorm:"VARCHAR(255)"`
	StateValueDesc string    `xorm:"VARCHAR(255)"`
	UpdatedAt      DateTime `xorm:"DATETIME"`
}
