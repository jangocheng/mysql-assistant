package mysqlutil

import (
	"fmt"
	"log"
	"owen2020/conn"

	"github.com/siddontang/go-mysql/replication"
	"gorm.io/gorm"
)

//GetGormConnByDbConfig 根据配置获取一个gorm db 句柄
func GetGormConnByDbConfig(DbConfig replication.BinlogSyncerConfig) *gorm.DB {
	return conn.GetGorm(DbConfig.Host, int(DbConfig.Port), DbConfig.User, DbConfig.Password)
}

func GetMysqlPosition(cfg replication.BinlogSyncerConfig) map[string]interface{} {
	sql := "show master status"
	results := make(map[string]interface{})

	query := GetGormConnByDbConfig(cfg)
	err := query.Raw(sql).Find(&results).Error
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// apputil.PrettyPrint(results)
	// fmt.Printf("%+v", results)
	return results
}

func GetDBs(gorm *gorm.DB) []string {
	sql := "show databases"

	result := []map[string]interface{}{}
	err := gorm.Raw(sql).Find(&result).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	dbs := []string{}
	// fmt.Println("dbs result")
	// apputil.PrettyPrint(result)

	for _, v := range result {
		if v["Database"] == "mysql" || v["Database"] == "information_schema" || v["Database"] == "performance_schema" || v["Database"] == "sys" {
			continue
		}
		dbs = append(dbs, v["Database"].(string))
	}

	// fmt.Println("dbs")
	// apputil.PrettyPrint(dbs)

	return dbs
}

func GetTableNames(gorm *gorm.DB, db string) []string {

	sql := "show tables from `" + db + "`"

	result := []map[string]interface{}{}

	err := gorm.Raw(sql).Find(&result).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// apputil.PrettyPrint(db)
	// apputil.PrettyPrint(result)

	tables := []string{}
	// fmt.Printf("%T", tables)

	for _, v := range result {
		tables = append(tables, v["Tables_in_"+db].(string))
	}

	return tables
}
