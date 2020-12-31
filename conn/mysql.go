package conn

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
	mysql "gorm.io/driver/mysql"
)

// time_zone时区问题 https://studygolang.com/articles/17313?fr=sidebar

// https://gorm.io/docs/index.html
//GetDefaultGorm
func GetDefaultGorm() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), 3306, os.Getenv("DB_DATABASE"), os.Getenv("DB_CHARSET"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

//GetGormWithLog
func GetGormWithLog() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), 3306, os.Getenv("DB_DATABASE"), os.Getenv("DB_CHARSET"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

//GetGormWithConfig 指定配置
func GetGormWithConfig(config *gorm.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), 3306, os.Getenv("DB_DATABASE"), os.Getenv("DB_CHARSET"))
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

//GetEventGorm
func GetEventGorm() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local", os.Getenv("DB_EVENT_USERNAME"), os.Getenv("DB_EVENT_PASSWORD"), os.Getenv("DB_EVENT_HOST"), 3306, os.Getenv("DB_EVENT_DATABASE"), os.Getenv("DB_EVENT_CHARSET"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetGorm(host string, port int, username string, password string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&loc=Local", username, password, host, 3306)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
