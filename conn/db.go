package conn

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

var DB *sql.DB

func InitDB(host string, port int, username string, password string) error {
	driverName := os.Getenv("DB_CONNECTION")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&loc=Local", username, password, host, 3306)

	DB, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3) // 3分钟最大保活，  取决于mysql server端的设置
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)

	return nil
}
