package apputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"owen2020/app/apputil/applog"

	"github.com/gin-gonic/gin"
)

func TryGet(c *gin.Context, k string) string {
	if c.Request.Method == "POST" || c.Request.Method == "PUT" {
		ret := GetFromBody(c, k)

		if ret != "" {
			return ret
		}
	}

	return c.Query(k)
}

func GetFromBody(c *gin.Context, k string) string {
	if c.ContentType() == "application/json" {
		ret := GetFromJson(c, k)
		retString := ret.(string)
		if retString != "" {
			return retString
		}
	}

	return c.PostForm(k)
}

func GetFromJson(c *gin.Context, k string) interface{} {
	parsed, _ := c.Get("jsonParsed")
	if "true" == parsed {
		return GetFromJmap(c, k)
	}

	// 把request的内容读取出来
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	fmt.Println(string(bodyBytes))

	// 把刚刚读出来的再写进去
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(bodyBytes, &jsonMap)
	c.Set("jsonParsed", "true")
	c.Set("jsonMap", jsonMap)
	if err != nil {
		applog.Logger.Error(err)
		return ""
	}

	return GetFromJmap(c, k)
}

func GetFromJmap(c *gin.Context, k string) interface{} {
	jsonMap, _ := c.Get("jsonMap")
	jmap := jsonMap.(map[string]interface{})
	v, has := jmap[k]
	if !has {
		return ""
	}
	return v
}
