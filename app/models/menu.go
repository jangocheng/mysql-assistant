package models

type Menu struct {
	MenuId    int      `json:"menu_id" xorm:"not null pk autoincr comment('{"name":"文档ID","desc":"哈哈哈哈哈哈","type":"password"}') INT(10)"`
	Title     string   `json:"title" xorm:"not null default '' comment('标题') VARCHAR(50)"`
	Pid       int      `json:"pid" xorm:"not null default 0 comment('{"name":"上级ID","desc":"","type":"select", "options":{"callback":"getMenuTree"}}') index INT(10)"`
	Sort      int      `json:"sort" xorm:"not null default 0 comment('排序（同级有效）') INT(10)"`
	Hide      int      `json:"hide" xorm:"not null default 0 comment('{"name":"是否隐藏","options":{"1":"否","2": "是"}}') TINYINT(1)"`
	Pathname  string   `json:"pathname" xorm:"comment('路由') VARCHAR(255)"`
	Iconfont  string   `json:"iconfont" xorm:"default '' comment('{"name":"图标"}') VARCHAR(255)"`
	Status    int      `json:"status" xorm:"default 1 TINYINT(4)"`
	CreatedAt DateTime `json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt DateTime `json:"updated_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	IsDeleted int      `json:"id_deleted" xorm:"default 0 TINYINT(4)"`
	DeletedAt DateTime `json:"deleted_at" xorm:"DATETIME"`
}
