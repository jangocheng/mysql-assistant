package conn

import (
	"database/sql"
	"fmt"
	"os"
	"owen2020/app/apputil"
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var SqlDB *sql.DB

/*
方式1 ， mysql链接池有database/sql 创建批定， gorm使用已有链接
此方式维护着项目的唯一的database/sql对象

项目中 调用sql.Open只调用一次
*/
func InitDB(host string, port int, username string, password string, database string) error {
	driverName := os.Getenv("DB_CONNECTION")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", username, password, host, port, database)

	DB, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3) // 3分钟最大保活，  取决于mysql server端的设置
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)

	apputil.PrettyPrint(DB)
	SqlDB = DB
	return nil
}

func DefaultEventGorm() *gorm.DB {
	gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: SqlDB}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return gdb
}

/**
方式2，  gorm创建链接，  配置gorm中包含的database/sql对象链接池信息
此方式维护着项目唯一的gorm对象

项目中 调用gorm.Open只调用一次
*/

var SyncerGormPool *gorm.DB

var EventGormPool *gorm.DB

func InitSyncerGormPool(host string, port int, username string, password string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&loc=Local", username, password, host, 3306)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)

	SyncerGormPool = db

	return nil
}

func GetSyncerGorm() *gorm.DB {
	return SyncerGormPool
}

func InitEventGormPool() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local", os.Getenv("DB_EVENT_USERNAME"), os.Getenv("DB_EVENT_PASSWORD"), os.Getenv("DB_EVENT_HOST"), 3306, os.Getenv("DB_EVENT_DATABASE"), os.Getenv("DB_EVENT_CHARSET"))
	mysqlDb := mysql.Open(dsn)

	db, err := gorm.Open(mysqlDb, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)

	EventGormPool = db

	return nil
}

//GetEventGorm
func GetEventGorm() *gorm.DB {
	return EventGormPool
}
