package models

type Administrator struct {
	AdministratorId int      `xorm:"not null pk autoincr INT(11)"`
	Username        string   `xorm:"VARCHAR(255)"`
	Password        string   `xorm:"VARCHAR(255)"`
	Name            string   `xorm:"VARCHAR(255)"`
	Avatar          string   `xorm:"comment('缩略图') VARCHAR(255)"`
	Status          int      `xorm:"TINYINT(1)"`
	CreatedAt       DateTime `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt       DateTime `xorm:"DATETIME"`
	IsDeleted       int      `xorm:"default 0 TINYINT(4)"`
	DeletedAt       DateTime `xorm:"DATETIME"`
}
