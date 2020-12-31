package apputil

import (
	"errors"
	"strconv"
)

const (
	MiddleError   = 600  // 600+ 中间件层
	ValidateError = 700  // 700+ 数据验证层不过通
	AppError      = 800  // 800+  控制器层逻辑和代码错误 业务代码
	LogicError    = 900  // 900+ 服务层和其他层 逻辑和代码错误
	MysqlError    = 1000 // 1000+ mysql层获取数据，存储数据错误
	CacheError    = 2000 // 2000+ redis等缓存层数据错误
	Other         = 9999 // 系统其他错误
)

var statusText = map[int]string{
	MiddleError:   "中间件验证不过通过",
	ValidateError: "数据验证不通过",
	AppError:      "系统出错",
	LogicError:    "系统逻辑错误",
	MysqlError:    "系统数据错误",
	CacheError:    "系统缓存错误",
	Other:         "系统其他错误",
}

// CommonErrorMap returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func CommonErrorMap(code int) string {
	errorMessage, ok := statusText[code]
	if ok {
		return errorMessage
	}
	return "未定义的错误码" + strconv.Itoa(code)
}

//MapedIntOrDefault 或得对应文本或默认值
func MapedIntOrDefault(emap map[int]string, code int) string {
	errorMessage, ok := emap[code]
	if ok {
		return errorMessage
	}

	return "未知" + strconv.Itoa(code)
}

func MapPhraseByNum(m *map[int]string, num int) (string, error) {
	phrase, has := (*m)[num]
	if has {
		return phrase, nil
	}
	return "", errors.New("not exsit")
}

func MapNumByPhrase(m *map[int]string, phrase string) (int, error) {
	for i, v := range *m {
		if v == phrase {
			return i, nil
		}
	}
	return 0, errors.New("not exsit")
}
