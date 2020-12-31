package msg

//ErrorMessage 自定义验证错误消息类型
type ErrorMessage map[string]string

//GetCustomErrorMsg 根据 field 和tag 获取自定义消息
func (e ErrorMessage) GetCustomErrorMsg(field string, tag string) (string, bool) {
	customErrorMessage, ok := e[field+"."+tag]
	if ok {
		return customErrorMessage, true
	}
	customErrorMessage, ok = e[field]
	if ok {
		return customErrorMessage, true
	}
	return "", false
}

//GetCustomErrorMsg 函数调用自定义map错误
func GetCustomErrorMsg(msg map[string]string, field string, tag string) (string, bool) {
	customErrorMessage, ok := msg[field+"."+tag]
	if ok {
		return customErrorMessage, true
	}
	customErrorMessage, ok = msg[field]
	if ok {
		return customErrorMessage, true
	}
	return "", false
}
