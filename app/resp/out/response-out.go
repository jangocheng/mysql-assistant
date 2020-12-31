package out

import (
	"net/http"
	"owen2020/app/apputil/applog"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code int    `json:"code" form:"code" xml:"code"`
	Msg  string `json:"msg" form:"msg" xml:"msg"`
}

type Response struct {
	Code int         `json:"code" form:"code" xml:"code"`
	Msg  string      `json:"msg" form:"msg" xml:"msg"`
	Data interface{} `json:"data" form:"data" xml:"data"`
}

//NewError 生成一个统一格式的错误信息
func NewError(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
	}
}

//NewErrorAndLog 生成一个统一格式的错误信息 记录日志
func NewErrorAndLog(code int, msg string) Response {
	applog.Logger.Error(msg)
	return Response{
		Code: code,
		Msg:  msg,
	}
}

//NewSuccess 生成一个统一格式的成功信息
func NewSuccess(data interface{}) Response {
	return Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	}
}

//NewErrorWithData 生成指定code和msg的错误信息， 是NewResponse的别名
func NewErrorWithData(code int, msg string, data interface{}) Response {
	return NewResponse(code, msg, data)
}

//NewSuccessManual 生成指定code和msg的成功信息， 是NewResponse的别名
func NewSuccessManual(code int, msg string, data interface{}) Response {
	return NewResponse(code, msg, data)
}

//NewResponse 生成格式统一的输出
func NewResponse(code int, msg string, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

//JSONOK 输出Json，状态码200
func (res Response) JSONOK(c *gin.Context) {
	c.JSON(http.StatusOK, res)
}

//XMLOK  输出xml，状态码200
func (res Response) XMLOK(c *gin.Context) {
	c.XML(http.StatusOK, res)
}

//HTMLOK 输出Html，状态码200
func (res Response) HTMLOK(c *gin.Context, name string) {
	c.HTML(http.StatusOK, name, res)
}

//JSON 输出Json，状态码自定义
func (res Response) JSON(c *gin.Context, status int) {
	c.JSON(status, res)
}

//XML 输出xml，状态码自定义
func (res Response) XML(c *gin.Context, status int) {
	c.XML(status, res)
}

//HTML 输出HTML，状态码自定义
func (res Response) HTML(c *gin.Context, name string, status int) {
	c.HTML(status, name, res)
}
