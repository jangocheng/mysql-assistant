package foundations

import "strconv"

// https://www.wenlincheng.com/posts/ef38.html
//对用户的姓名、电话号码、邮箱、身份证等信息进行脱敏处理。

//HidePhone 手机号脱敏
func HidePhone(phone string) string {
	rs := []rune(phone)
	result := string(rs[0:3]) + "****" + string(rs[7:11])
	return result
}

//HidePhoneInt 手机号脱敏
func HidePhoneInt(phone int) string {
	return HidePhone(strconv.Itoa(phone))
}
