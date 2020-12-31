package request

import (
	"owen2020/app/apputil"
	"owen2020/app/resp/out"

	"github.com/gin-gonic/gin"
)

type login struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	KeepLogin string
	JumpURL   string
}

//DefaultJumpURL 登录成功默认跳转页
var DefaultJumpURL string = "/member/userinfo_detail.html"

//GetLoginParams 获取登录入参
func GetLoginParams(c *gin.Context) (*login, error) {
	loginParams := &login{}

	err := c.ShouldBind(&loginParams)
	if nil != err {
		out.NewError(apputil.ValidateError, "验证失败:"+err.Error()).JSONOK(c)
		return nil, err
	}

	return loginParams, nil
}

//SetDefault  登录默认跳转页面
func SetDefault(loginParams *login) {
	if loginParams.JumpURL == "" {
		loginParams.JumpURL = DefaultJumpURL
	}
}
