package admin

import (
	"fmt"
	"owen2020/app/resp/out"

	"github.com/gin-gonic/gin"
)

//Test try it
func Test(c *gin.Context) {
	fmt.Printf("%+v", c.Request)
	out.NewSuccess(c.Request).JSONOK(c)
}
