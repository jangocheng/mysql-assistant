package models

import (
	"time"
)

type State struct {
	CreatedAt      time.Time `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	DeletedAt      time.Time `xorm:"DATETIME"`
	IsDeleted      int       `xorm:"default 0 TINYINT(4)"`
	StateClassId   int       `xorm:"INT(11)"`
	StateId        int       `xorm:"not null pk autoincr INT(11)"`
	StateValue     string    `xorm:"VARCHAR(255)"`
	StateValueDesc string    `xorm:"VARCHAR(255)"`
	UpdatedAt      time.Time `xorm:"DATETIME"`
}
