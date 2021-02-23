package middleware

import (
	"errors"
	"fmt"
	"os"
	"owen2020/app/apputil/applog"
	"owen2020/app/resp/out"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//AdminToken 如果是已登录用户， 从token中提取用户ID
func AdminToken(c *gin.Context) {
	err := AdminCheckToken(c)
	if err != nil {
		c.Abort()
	}
	c.Next()
}

//AdminCheckToken 如果是已登录用户， 从token中提取用户ID
// 提取这个方法是为NoRoute中REST使用
func AdminCheckToken(c *gin.Context) error {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		tokenString, _ = c.Cookie("AdminAuthorization")
	}
	if tokenString == "" {
		out.NewError(600, "缺失认证token").JSONOK(c)
		return errors.New("缺失认证token")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSampleSecret := os.Getenv("ADMIN_JWT_SECRET")
		fmt.Println(hmacSampleSecret)
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSampleSecret), nil
	})

	if nil != err {
		applog.Logger.WithFields(log.Fields{"token": tokenString, "error": err.Error()}).Info("admin-token解析token失败")
		out.NewError(600, "admin-token解析token失败").JSONOK(c)
		return errors.New("admin-token解析token失败")
	}

	if !token.Valid {
		applog.Logger.WithFields(log.Fields{"token": tokenString, "secret": os.Getenv("JWT_SECRET")}).Info("admin-token解析无效")
		out.NewError(600, "admin-token解析无效").JSONOK(c)
		return errors.New("admin-token解析无效")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		applog.Logger.WithFields(log.Fields{"claims": claims}).Info("admin-token claims解析无效")
		out.NewError(600, "admin-token, claims无效").JSONOK(c)
		return errors.New("admin-token, claims无效")
	}
	c.Set("admin_id", claims["uid"])

	return nil
}
