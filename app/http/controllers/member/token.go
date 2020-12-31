package member

import (
	"os"
	"owen2020/app/apputil/foundations"
	"owen2020/app/http/controllers/member/request"
	"owen2020/app/models/dao"
	"owen2020/app/resp/out"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

func GetTokenList(c *gin.Context) {

	return
}

func GetTokenInfo(c *gin.Context) {
	return
}

// TokenAdd 用户登录
// @Summary 根据手机号密码登录
// @Accept  json
// @Produce  json
// @Param phone body string true "手机号用户名"
// @Param password body string true "密码"
// @Response 200 {object} apputil.Response
// @Router /member/v1/token [post]
func TokenAdd(c *gin.Context) {
	loginParams, err := request.GetLoginParams(c)
	if nil != err {
		return
	}
	request.SetDefault(loginParams)

	memberInfo, err := dao.GetMemberInfoByMobile(loginParams.Username)
	if nil != err {
		out.NewError(800, "用户名或密码错误").JSONOK(c)
		return
	}

	// 密码不正确
	genPwd := foundations.EncryptWord(loginParams.Password, nil)
	if memberInfo.Password != genPwd {
		out.NewError(800, "用户名或密码错误2").JSONOK(c)
		return
	}

	tokenString, err := genMemberToken(strconv.Itoa(memberInfo.MemberId))
	if err != nil {
		out.NewError(800, "生成jwt出错").JSONOK(c)
	}

	// 设置cookie
	c.SetCookie("Authorization", tokenString, 1000, "/", c.Request.Host, false, true)

	// 如果是ajax，不需要跳转，直接输出结果
	if strings.EqualFold(c.Request.Header.Get("X-Requested-With"), "XMLHttpRequest") {
		out.NewSuccess(gin.H{"token": tokenString, "info": gin.H{"member_id": memberInfo.MemberId}, "jump_url": loginParams.JumpURL}).JSONOK(c)
		return
	}
	// 登录完成跑转到jump_url
	c.Redirect(302, loginParams.JumpURL)
	return
}

//genMemberToken
func genMemberToken(memberId string) (string, error) {
	hmacSampleSecret := os.Getenv("JWT_SECRET")
	// 生成token https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"uid": memberId, "exp": time.Now().Unix() + 86400*30})

	return jwtToken.SignedString([]byte(hmacSampleSecret))
}

func TokenEdit(c *gin.Context) {
	return
}

//TokenDelete 退出登录
func TokenDelete(c *gin.Context) {
	// 设置cookie
	c.SetCookie("Authorization", "", -1, "/", c.Request.Host, false, true)

	out.NewSuccess(gin.H{"jump_url": ""}).JSONOK(c)
}
