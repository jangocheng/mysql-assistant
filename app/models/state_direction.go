package models

import (
	"time"
)

type StateDirection struct {
	CreatedAt        time.Time `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	DeletedAt        time.Time `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	IsDeleted        int       `xorm:"default 0 TINYINT(4)"`
	StateClassId     int       `xorm:"default 0 INT(11)"`
	StateDirectionId int       `xorm:"not null pk autoincr INT(11)"`
	StateFrom        int       `xorm:"default 0 INT(11)"`
	StateTo          int       `xorm:"default 0 INT(11)"`
	UpdatedAt        time.Time `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
}
