package foundations

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"owen2020/app/apputil/applog"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenToken(uid int) (string, error) {
	hmacSampleSecret := os.Getenv("JWT_SECRET")

	// 生成token https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"uid": uid, "exp": time.Now().Unix() + 86400*30})

	return jwtToken.SignedString([]byte(hmacSampleSecret))
}

func VerifyGetToken(s string) (*jwt.Token, error) {
	//token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
	//	// Don't forget to validate the alg is what you expect:
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	//	}
	//
	//	hmacSampleSecret := os.Getenv("JWT_SECRET")
	//	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	//	return []byte(hmacSampleSecret), nil
	//})

	token, err := jwt.Parse(s, keyFuncX)

	if nil != err {
		applog.Logger.WithFields(log.Fields{"token": s}).Info("解析token失败")
		return nil, errors.New("解析token失败")
	}

	if !token.Valid {
		applog.Logger.WithFields(log.Fields{"token": s, "secret": os.Getenv("JWT_SECRET")}).Info("解析无效")
		return nil, errors.New("解析无效")
	}

	return token, nil
}

func keyFuncX(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	hmacSampleSecret := os.Getenv("JWT_SECRET")
	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte(hmacSampleSecret), nil
}