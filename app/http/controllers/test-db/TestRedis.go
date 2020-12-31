package test_db

import (
	"context"
	"fmt"
	"net/http"
	"owen2020/conn"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func Redigo(c *gin.Context) {
	redisConn, _ := conn.GetRedisConn()
	defer redisConn.Close()

	setReply, err := redisConn.Do("SET", "test-redis", "yes")

	if nil != err {
		fmt.Println(err.Error())
	}
	fmt.Println(setReply)

	getReply, _ := redisConn.Do("GET", "test-redis")

	fmt.Println(getReply)

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": string(getReply.([]byte))})

}

func RedisRaw(c *gin.Context) {
	rdb := conn.GetRawRedis()
	defer rdb.Close()

	ctx := context.Background()

	setRet, err := rdb.Set(ctx, "raw-key", "hahaha", time.Minute*5).Result()
	if err != nil {
		fmt.Println("raw redis set err:", err.Error())
	}
	fmt.Println("set result:", setRet)

	val, err := rdb.Get(ctx, "raw-key").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	fmt.Println(val)

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": val})

}
