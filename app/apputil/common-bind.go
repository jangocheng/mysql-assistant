package apputil

import (
	"owen2020/app/resp/out"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

//ShouldBindOrError 绑定数据
func ShouldBindOrError(c *gin.Context, model interface{}) error {
	// 中文错误输出转换器
	zh := zh.New()
	// 统一万能转换器入口， 它是管理转换器的，可以同时包含中文，英文等多个转换器
	uni := ut.New(zh)
	// 从统一万能转换器获得，包装了zh中文转换器的translator
	trans, _ := uni.GetTranslator("zh")
	// zh_translations.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans)
	// 为bind中的validator注册trans
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zh_translations.RegisterDefaultTranslations(v, trans)
	}

	// err := c.Bind(&activity) // Bind 会输出400header头
	// 或者需要自定义错误，有下面的代码
	err := c.ShouldBind(model) // ShouldBind不输出400header头 ,就这个区别
	if nil != err {
		out.NewErrorWithData(ValidateError, CommonErrorMap(ValidateError), err.(validator.ValidationErrors).Translate(trans)).JSONOK(c)
		return err
	}

	return nil
}
