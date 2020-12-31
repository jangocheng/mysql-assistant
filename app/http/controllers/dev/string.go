package dev

import (
	"owen2020/app/apputil"
"owen2020/app/resp/out"
	"strings"

	"github.com/gin-gonic/gin"
)

func TestString(c *gin.Context) {
	stringX := "admin/v1/hahah/asfsf/bbb/dd"

	stringSlice := strings.SplitN(stringX, "/", 3)

	out.NewSuccess(stringSlice).JSONOK(c)
}

func TryGet(c *gin.Context) {
	ret := apputil.TryGet(c, "aaa")

	out.NewSuccess(ret).JSONOK(c)
}
