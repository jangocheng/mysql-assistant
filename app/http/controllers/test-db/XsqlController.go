package test_db

// http://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Cmysql/select%E6%93%8D%E4%BD%9C.html
// https://godoc.org/github.com/jmoiron/sqlx
// sqlx的数据库操作学习
// 对比原生sql 直接读取到结构体中，真的太友好了。
import (
	"fmt"
	"net/http"
	"owen2020/app/models"
	"owen2020/app/resp/out"
	"owen2020/conn"

	"github.com/gin-gonic/gin"
)

type RowData []interface{}

func SqlxGetList(c *gin.Context) {
	sqlx := conn.GetSQLx()

	var person []Person

	err := sqlx.Select(&person, "select * from `person`")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": person})
}

func SqlxGetInfo(c *gin.Context) {
	sqlx := conn.GetSQLx()
	var person models.Person

	err := sqlx.Get(&person, "select * from `person`")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": person})
}

func SqlxInsert(c *gin.Context) {
	// 使用Exec方法， 和原生操作一样， 没有什么区
	c.JSON(http.StatusOK, gin.H{"msg": "使用Exec方法插入，和原生sql操作一样， 没有封装什么"})
}

func SqlxUpdate(c *gin.Context) {
	// 使用Exec方法， 和原生操作一样， 没有什么区
	c.JSON(http.StatusOK, gin.H{"msg": "使用Exec方法更新,和原生sql操作一样， 没有封装什么"})
}

func SqlxDelete(c *gin.Context) {
	// 使用Exec方法， 和原生操作一样， 没有什么区
	c.JSON(http.StatusOK, gin.H{"msg": "使用Exec方法删除,和原生sql操作一样， 没有封装什么"})
}

func SqlxGetListMap(c *gin.Context) {
	sqlx := conn.GetSQLx()

	rows, err := sqlx.Queryx("select * from `person`")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var lists []RowData
	// for rows.Next() {
	// 	row := make(map[string]interface{})

	// 	// sli := make([]interface{}, 10)

	// 	// err := rows.MapScan(row)
	// 	sli, err := rows.SliceScan()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	lists = append(lists, row)
	// }

	for rows.Next() {
		// sli, err := rows.SliceScan()
		// if err != nil {
		// 	out.NewError(700, err.Error()).JSONOK(c)
		// 	return
		// }
		sli, _ := rows.SliceScan()
		// fmt.Printf("%+v", sli)
		fmt.Printf("%s", sli)
		lists = append(lists, sli)
	}

	out.NewSuccess(lists).JSONOK(c)
}
