package admin

import (
	"github.com/gin-gonic/gin"
	"owen2020/app/models/dao"
	"owen2020/app/resp/out"
)

func GetStateClassList(c *gin.Context) {
	list := dao.GetMenuList()
	ml := menuList(list)
	//ret := processToTree(list, 0, 0)
	ret := ml.processToTree(0, 0)
	out.NewSuccess(ret).JSONOK(c)
}
