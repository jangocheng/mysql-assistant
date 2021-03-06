package admin

import (
	"owen2020/app/apputil"
	"owen2020/app/apputil/foundations"
	"owen2020/app/models/dao"
	"owen2020/app/resp/out"
	"strings"

	"github.com/gin-gonic/gin"
)

type login struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	KeepLogin string
	JumpURL   string `json:"jump_url" form:"jump_url"`
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


func CreateToken(c *gin.Context) {
	loginParams, err := GetLoginParams(c)
	if nil != err {
		return
	}

	adminInfo := dao.GetAdminInfo(loginParams.Username)
	if nil == adminInfo {
		out.NewError(800, "用户名或密码错误").JSONOK(c)
		return
	}

	// 密码校验
	genPwd := foundations.EncryptWord(loginParams.Password, nil)
	if adminInfo.Password != genPwd {
		out.NewError(800, "用户名或密码错误2").JSONOK(c)
		return
	}

	tokenString, err := foundations.GenToken(adminInfo.AdministratorId)
	if err != nil {
		out.NewError(800, "生成jwt出错").JSONOK(c)
	}

	// 设置cookie
	c.SetCookie("AdminAuthorization", tokenString, 86400, "/", c.Request.Host, false, true)

	// 如果是ajax，不需要跳转，直接输出结果
	if strings.EqualFold(c.Request.Header.Get("X-Requested-With"), "XMLHttpRequest") {
		out.NewSuccess(gin.H{"token": tokenString, "info": gin.H{"admin_id": adminInfo.AdministratorId}, "jump_url": loginParams.JumpURL}).JSONOK(c)
		return
	}
	// 登录完成跑转到jump_url
	c.Redirect(302, loginParams.JumpURL)
	return
}
