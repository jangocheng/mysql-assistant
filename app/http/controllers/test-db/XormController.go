package test_db

// https://gobook.io/read/gitea.com/xorm/manual-zh-CN/ 中文文档
// https://github.com/go-xorm/xorm

import (
	"log"
	"net/http"
	"owen2020/conn"
	"owen2020/app/models"

	"github.com/gin-gonic/gin"
)

func XormGetList(c *gin.Context) {
	engine := conn.GetXormWithLog()
	// results, err := engine.Query("select * from person")
	results, err := engine.QueryString("select * from person")
	if nil != err {
		log.Println("err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": results})
}

func XormGetListFind(c *gin.Context) {
	engine := conn.GetXorm()
	var personList []Person
	// results, err := engine.Query("select * from person")
	err := engine.Table("person").Where("user_id < ?", 10).Limit(10, 0).Find(&personList)
	if nil != err {
		log.Println("err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": personList})
}

func XormGetInfo(c *gin.Context) {
	engine := conn.GetXorm()
	var personInfo models.Person

	has, err := engine.Table("person").Where("user_id = ?", 3).Get(&personInfo)
	if nil != err {
		log.Println("err", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if false == has {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "没有相关记录，查了个寂寞"})
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": personInfo})
}

func XormInsert(c *gin.Context) {
	personInfo := &Person{
		0,
		"124",
		1,
		"owe@ss",
		"",
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": personInfo})

}

// XormInsertFillId 插入填充必读：https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-04/index.html
func XormInsertFillId(c *gin.Context) {

}

func XormUpdate(c *gin.Context) {

}

func XormDelete(c *gin.Context) {

}
