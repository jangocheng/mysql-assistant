package test_db

//@see https://github.com/go-sql-driver/mysql
// 原生sql的写法
// go-sql-driver/mysql实现了"database/sql"
// 原生的写法 增删改查学习
// https://blog.csdn.net/westhod/article/details/80799266 go语言数据库查询后对结果的处理方法的探讨
// 最后的取值比较蛋疼，scan取值要一个一个的指定字段
// 用完原生sql的各种功能， 个人认为是读取数据最为费劲
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"owen2020/conn"

	"github.com/gin-gonic/gin"
)

// https://blog.csdn.net/qq_33679504/article/details/100533703
// json tags指定类型只支持 string ,number, boolean 不支持int, 因为以接口输出就这些类型了。
// 明天需要尝试一下到底是否支持超大数字
// https://cloud.tencent.com/developer/section/1141542  json数据输出格式说明
// bool, for JSON booleans
// float64, for JSON numbers
// string, for JSON strings
// []interface{}, for JSON arrays
// map[string]interface{}, for JSON objects
// nil for JSON null
type Person struct {
	UserId    int    `json:"user_id" db:"user_id"`
	Username  string `json:"user_name" db:"username"`
	Sex       int    `json:"sex,string"`
	Email     string `json:"email" db:"email"`
	CreatedAt string `json:"created_at" db:"created_at" time_format:"2006-01-02 15:04:05" time_utc:"8"`
	// CreatedAt []byte `db:"created_at" time_format:"2006-01-02 15:04:05" time_utc:"8"`
}

//RawSqlGetList 列表
func RawSqlGetList(c *gin.Context) {
	// rawDb := conn.GetRawDb()

	// rows, err := rawDb.Query("select * from `person`")
	// if nil != err {
	// 	log.Println("查询失败:", err.Error())
	// }
	// defer rows.Close()

	// println(rows)

	// c.JSON(http.StatusOK, gin.H{"rows": rows})
	getList1(c)
	// getList2(c)
}

// getList
// rows.Next的作用有两，1是检查是否还有数据， 2是移动指针，未移动指针时rows.scan内容为空。 它默认在“顶点”，而不是第一行。
func getList1(c *gin.Context) {
	rawDb := conn.GetRawDb()

	rows, err := rawDb.Query("select * from `person`")
	if nil != err {
		log.Println("查询失败:", err.Error())
	}
	defer rows.Close()

	rows.Next()
	var person Person
	rows.Scan(&person.UserId, &person.Username, &person.Sex, &person.Email, &person.CreatedAt)
	fmt.Printf("%+v\n", person)
	c.JSON(http.StatusOK, gin.H{"rows": person})
	return
}

func getList2(c *gin.Context) {
	rawDb := conn.GetRawDb()

	rows, err := rawDb.Query("select * from `person`")
	if nil != err {
		log.Println("查询失败:", err.Error())
	}
	defer rows.Close()

	// 获取列名
	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	rows.Next()
	rows.Scan(scanArgs...)

	rowMap := make(map[string]string)
	var value string
	for i, col := range values {
		// Here we can check if the value is nil (NULL value)
		if col != nil {
			value = string(col)
			rowMap[columns[i]] = value
		}
	}

	fmt.Printf("%+v\n", rowMap)
	fmt.Printf("%+v\n", values)
	c.JSON(http.StatusOK, gin.H{"rows": rowMap, "values": values, "columns": columns})

}

//RawSqlGetInfo 详情
// getInfo和getList没什么差别， 重点在于可以有个QueryRow, 直接返回一行。
// 只是还没搞懂QueryContext和Query的区别
// https://www.reddit.com/r/golang/comments/ffvijc/databasesql_query_vs_querycontext/ 可以取消读取sql, 有此情况没有必要再次读取
func RawSqlGetInfo(c *gin.Context) {

}

//RawSqlInsert 插入
func RawSqlInsert(c *gin.Context) {
	rawDb := conn.GetRawDb()
	ret, err := rawDb.Exec("insert into `person` (`username`, `sex`, `email`) values(?, ?, ?)", "testInsert", 1, "owen@libu")
	if nil != err {
		log.Print("插入数据错误", err.Error())
	}

	rows, err := ret.RowsAffected()
	if 0 == rows {
		log.Print("插入数据错误,影响行数为0", err.Error())
	}

	id, _ := ret.LastInsertId()

	c.JSON(http.StatusOK, gin.H{"msg": "插入成功", "id": id})
}

//RawSqlUpdate 更新
func RawSqlUpdate(c *gin.Context) {
	rawDb := conn.GetRawDb()
	ret, err := rawDb.Exec("update `person` set username=? where user_id=?", "testUpdate", 3)
	if nil != err {
		log.Print("更新数据错误", err.Error())
	}
	rows, err := ret.RowsAffected()
	if 0 == rows {
		log.Print("插入数据错误,影响行数为0", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"msg": "更新成功", "affect_row": rows})
}

//RawSqlDelete 删除
func RawSqlDelete(c *gin.Context) {
	rawDb := conn.GetRawDb()
	ret, err := rawDb.Exec("delete from `person` where user_id=?", 4)
	if nil != err {
		log.Print("删除数据错误", err.Error())
		return
	}
	rows, err := ret.RowsAffected()
	if 0 == rows {
		log.Print("删除数据错误,影响行数为0", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"msg": "删除成功", "affect_row": rows})
}
