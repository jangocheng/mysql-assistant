package handle_binlog

import (
	"fmt"
	"os"
	"owen2020/cmd/command/mysqlutil"
	"regexp"
	"strings"
)

func InitDBTables() {
	gorm := mysqlutil.GetGormConnByDbConfig(Cfg)

	dbs := mysqlutil.GetDBs(gorm)

	for _, db := range dbs {
		FlushDBTables(db)
	}
}

func FlushDBTables(db string) {
	tables := mysqlutil.GetTableNames(mysqlutil.GetGormConnByDbConfig(Cfg), db)
	for _, v := range tables {
		FlushTableIdentifierNameMap(db, v)
	}
}

// https://dev.mysql.com/doc/refman/8.0/en/identifiers.html
func FlushTableIdentifierNameMap(db string, table string) {
	ok := FilterTable(db, table)
	if !ok {
		return
	}

	sql := "show full columns from `" + table + "` from `" + db + "`"
	// sql := "show full FIELDS from `user` from `codeper`"

	results := []map[string]interface{}{}

	gorm := mysqlutil.GetGormConnByDbConfig(Cfg)

	err := gorm.Raw(sql).Find(&results).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	// apputil.PrettyPrint(results)
	columns := TableColumnIdentify{}
	for i, v := range results {
		columns[i] = v["Field"].(string)
	}

	DBTables[db+"."+table] = columns
}

//FilterTable 检查table 是否在正则中
func FilterTable(db string, table string) bool {
	if db == os.Getenv("DB_EVENT_DATABASE") {
		return false
	}

	filters := strings.Split(Filter, ",")
	for _, v := range filters {
		ok, _ := regexp.Match(v, []byte(db+"."+table))
		if ok {
			return true
		}
	}

	return false
}
