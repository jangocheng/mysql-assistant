package dev

import (
	"owen2020/app/resp/out"

	"github.com/gin-gonic/gin"
)

//JsonToken https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func JsonToken(c *gin.Context) {

	// tokenString := c.GetHeader("Authorization")
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDkxMzY0MTIsInVpZCI6IjEifQ.sNgY8ntuoYxswSCBHfSY74VJac4ST2axj1S7YnJCouE"

	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// Don't forget to validate the alg is what you expect:
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	hmacSampleSecret := os.Getenv("JWT_SECRET")
	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	// 	return []byte(hmacSampleSecret), nil
	// })

	// if nil != err {
	// 	fmt.Println("解析token失败")
	// 	return
	// }

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	fmt.Println(claims)
	// } else {
	// 	fmt.Println(err)
	// }

	memberID, _ := c.Get("member_id")

	out.NewSuccess(gin.H{"member_id": memberID}).JSONOK(c)
}
