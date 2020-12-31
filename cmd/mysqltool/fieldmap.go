package mysqltool

import (
	"log"
	"owen2020/conn"
)

type FieldInfo struct {
	Name    string
	Primary bool
}

//https://dev.mysql.com/doc/refman/5.7/en/information-schema-columns-table.html

//https://dev.mysql.com/doc/refman/8.0/en/show-columns.html
// type FieldMap map[int]string

//GetFieldMap 根据db 和 table name获取字段对应
func GetFieldMap(db string, table string) map[int]string {
	sql := "show EXTENDED FULL FIELDS from `" + table + "` from `" + db + "`"
	// sql := "show FULL COLUMNS from `" + table + "` from `" + db + "`"

	engine := conn.GetXormWithLog()
	results, err := engine.QueryString(sql)
	if err != nil {
		log.Fatal(err)
	}
	var m map[int]string

	var i int = 0
	for _, v := range results[0] {
		m[i] = v
		i++
	}

	return m
}
