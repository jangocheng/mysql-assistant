# golang学习

## 优秀项目
https://blog.csdn.net/xx666zz/article/details/89337889

# 已完成学习
 - route分组与配置
 - route - restful接口配置
 - .env文件使用，读取配置
 - controller, model, config, connection分层
 - raw sql操作学习完成
 - sqlx操作学习完成
 - xorm操作学习完成
 - gorm操作学习完成
 - raw redis操作学习完成
 - redigo 操作学习完成
 - oss上传 完成
 - bind 学习完成， 需要再熟悉tags相关参数
 - validator 学习完成， 需要再熟悉tag相关参数
 - validator 自定义错误语输出完成 和i18n完成
 - 服务热重启完成
 - cookie操作完成
 - swagger自动生成文档 https://github.com/swaggo/gin-swagger， https://swaggo.github.io/swaggo.io/declarative_comments_format/
 - cmd 命令行配置完成
 - grpc server ,gprc client 调试完成
 - canal-client 消费binlog完成



# 项目主要功能模块
 - cmd 命令行
 - http  web
 - config 配置
 - conn mysql,redis,oss等链接
 - routes 路由
 - storage 临时文件，日志等 

# 项目主要分层
 - reqt 接口入参结构体
 - reqt/msg  入参校验错误提示语
 - model 对应数据库结构体
 - resp 接口出参， 对model层的status等字段 处理文字输出
 - models/dao 一些快捷方法， 使用gorm， raw sql 等获取数据
 - controller 控制器
 - service 复杂逻辑 可由controller层提取出来
 - grpc 和 proto只是实验性质， 和model层不能完美整合一体服务， 所以学习调试完成就OK了。

## xorm/cmd自动生成model文件
安装失败的两种情况：  
1： 未生成xorm命令到bin目录中：  
解决办法： 添加GOPATH  GOROOT到环境变量  

2： 未go get 安装时提示错误：  
解决办法： 清除 gopkg目录下的go-xorm库包， 重新安装（因为版本问题。）

#### 根据table生成model命令
xorm reverse "mysql" "username:passsword@tcp(host:3306)/dbname?charset=utf8" "/Users/owen/go/pkg/mod/github.com/go-xorm/cmd/xorm@v0.0.0-20190426080617-f87981e709a1/templates/goxorm" 


# todo
- 数据库type字段 对应含义
- 原生mysqlbinlog 解析事件， 这个可以完的话就完美了。
- container/* 学习使用
- zipkin log封装



原来是这样，如果datetime格式就不要parseTime选项，以string输出就可以
加上parseTime选项 datetime会自动转换成time.Time项。
而time.Time默认format模式出输 不是“2006-01-02 15:04:05”，  默认输出类型："2020-11-18T21:56:22+08:00" 就是这个样子。
所以，如果加了parseTime参数， 要想办法设置time_format选项。


# 验证器自定义错误内容，中文，英文trans
验证器输出中文已完成。
记录一下输出中文错误提示的方式。  

```
    // 中文错误输出转换器
	// zh := zh.New()
	// 统一万能转换器入口， 它是管理转换器的，可以同时包含中文，英文等多个转换器
	uni := ut.New(zh.New())

	// 从统一万能转换器获得，包装了zh中文转换器的translator
	trans, _ := uni.GetTranslator("zh")
    zh_translations.RegisterDefaultTranslations(validate, trans)
```

输出自定义错误，已实现
```
定义在：github.com/go-playground/validator/v10@v10.2.0/errors.go
validator.Struct验证返回是 type ValidationErrors []FieldError
FieldError 可能获取field,tag和value等字段, 所以可以定义map错误提示语，自定义错误输出

// 实现自定义错误输出
errMessage := make(map[string]string)
errMessage["ActivityPic"] = "活动图片不存在"
customErr, ok := errMessage[err.Field()]
if ok {
    c.JSON(http.StatusOK, gin.H{"code": 600, "msg": customErr})
    return
} else {
    c.JSON(http.StatusOK, gin.H{"code": 600, "msg": errs.(validator.ValidationErrors).Translate(trans)})
    return
}

```

## bind注册trans， zh_translations.RegisterDefaultTranslations只能调用一次，不能既注册到bind已注册到new出来的validator，原理不明。
想给new出来的validator注册，估计得copy一个trans出来
```
	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"

    // 为bind中的validator注册trans
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zh_translations.RegisterDefaultTranslations(v, trans)
	}
```


# 热重启 服务软重启   成功-完成, 这个文章是必读的， socket复用。
https://blog.csdn.net/tomatomas/article/details/94839857  

要执行 kill -1 , kill -2  
保证新编译的包替换原包， 需要是编译过的，go run直接跑的不能重启
因为go run 也是编译后运行，编译到临时目录中，如：/var/folders/pt/2rhc9t5x1gl8108gfgs_pn800000gn/T/go-build728960581/b001/exe/main   
重新执行go run 生成的main可执行文件，是在新的临时目录。  


# validator 自定义错误输出

这个文章介绍了如何处理err， 一步步引导，很详细
> https://blog.depa.do/post/gin-validation-errors-handling#toc_3  
> https://github.com/go-playground/validator/blob/master/_examples/translations/main.go



# 参考
- https://riptutorial.com/go/example/29299/restfull-projects-api-with-gin

- https://dev.to/cathysmith/how-to-build-restful-api-service-using-gin-gonic-framework-1m18

# 根据table自动生成struct有很多工具，重点是生成tags带着什么名称
如用 xorm生成的就带有xorm标识，其他orm类不一支持，可能需要自已开发。  
所以，既然选用xorm，那用就xorm的reverse生成struct，不用考虑其他工具。  
https://segmentfault.com/a/1190000017090805  


# 关于i18n的解释
>https://baike.baidu.com/item/I18N/6771940?fr=aladdin
l18n 为internationalization  国际化， 说白了就是多语言显示， 如切换中文，英文，其他文种等。  
和k8s一样，是一个长单词的缩写。  


# gin链路日志, 这个需要好好看一下， 链路日志，统一封装输出 和 locale 多语言使用
https://www.mdeditor.tw/pl/2ez0

# go 使redis， 
> https://blog.csdn.net/han0373/article/details/80611111
> 使用推荐的https://godoc.org/github.com/garyburd/redigo/redis


# 关于go包的说明，和包中变量做用域的说明
> https://studygolang.com/articles/16765?fr=sidebar


# 关于什么时候可以用nil的问题
> https://studygolang.com/articles/19625?fr=sidebar
```
指针，函数，interface，slice，channel和map的零值都是nil

以上几个类型都可为返回时都可以为nil， 那error是 interface ,所以可以赋值为nil

https://blog.csdn.net/whatday/article/details/107992618  error使用说明
```

# 指针变量的使用， 函数内使用指针 可返回nil
> https://www.cnblogs.com/ywjfx/p/10386928.html
> @see app/http/controllers/activity/activity.go:90:20


# go tool pprof打开方式由sublime变更为chrome
> https://blog.csdn.net/jiaolongdy/article/details/50945684

# golang 中文手册
>https://studygolang.com/pkgdoc

# golang英文手册
> https://blog.golang.org/pprof

# go web项目目录规划案例
>https://goframe.org/start/index

# log 日志学习 封装框架日志
> https://www.cnblogs.com/rickiyang/p/11074164.html

# go操作写文件flag标识说明
> https://studygolang.com/articles/25372?fr=sidebar

# go SetCookie时遇到的新问题， URL解析，获得域名，如果golang和nginx搭配 怎么获取hostname等
解决golang获取不到nginx配置的域名(www.xxx.com:18801/testapi/test)问题
>https://www.cnblogs.com/dfsxh/articles/10430292.html
Uri说明
> https://blog.csdn.net/g777520/article/details/79700368
Golang URL解析
>https://studygolang.com/articles/10050

# 用golang fastcgi与nginx配合写web
https://blog.csdn.net/ijibu/article/details/11975615


# 哪些想要的功能是因为中了PHP的毒
 - 根据函数名字（字符串）去调用一个函数。
 - 判断对象中方法是否存在。
上面两个功能都能通过反射实现。 但是，，，，，   
go是编译性静态语言！！！！ 就应该避免有这种想法！！！！  
非要实现的话，也可以，如下：
> https://mrwaggel.be/post/golang-reflect-if-initialized-struct-has-member-method-or-fields/


# golang数据竞态, 两个线程操作同一个变量的问题。
https://blog.csdn.net/love666666shen/article/details/108648365


# mysql中的time.Time 输出成 2006-01-02 15:04:05 格式， 也就是自定义格式和encode/json包的搭配使用， 这个需要学习encode/json包

golang结构体json的时间格式化解决方案
> https://blog.csdn.net/weixin_34220179/article/details/89837776

golang的json的时间格式化解决方案
> https://studygolang.com/articles/8962

golang 自定义time.Time json输出格式
>https://www.cnblogs.com/xiaofengshuyu/p/5664654.html

第三个文章是最好的， 因为除了json输出需要的MarshalJSON， 还有string需要的String()。  在记录日志， fmt.Print等函数，输出都不一定是json. 所以，String也是要实现的， 方便阅读。
> https://blog.csdn.net/sunansheng/article/details/89632573

> golang的fmt包String(),Error(),Format(),GoString()的接口实现
https://blog.csdn.net/weixin_34064653/article/details/91439415?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromBaidu-2.control&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromBaidu-2.control


# 关于mysql类型转换和 datetime存在问题
> https://dev.mysql.com/doc/refman/5.6/en/using-date.html

# 项目接口
> http://120.48.27.136:8000/

# 文档接口, 还是用Yapi
> https://www.jianshu.com/p/1761449ddf2b
> http://120.48.27.136:3000/project/11/interface/api

docgen 可以生成无返回内容的文档。
> https://github.com/thedevsaddam/docgen
```
docgen server -f go-activity.postman_collection.json -p 7000
```



# golang 打开本地浏览器
>https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang

# 函数传递slice是指针， 但在函数内后， 会得到一个新的slice指钍。
>https://blog.csdn.net/u013536232/article/details/105547626

# how-to-append-to-a-slice-pointer-receiver
> https://stackoverflow.com/questions/36854408/how-to-append-to-a-slice-pointer-receiver


#  imaginary part of complex numbers   iota的全称
> https://www.jb51.cc/go/186938.html, https://ask.csdn.net/questions/1011576
> https://srfi.schemers.org/srfi-1/srfi-1.html#iota

# iota完全理解
> https://studygolang.com/articles/22468?fr=sidebar

# coast指定数据类型, 数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
> https://zhuanlan.zhihu.com/p/137489817


# 理解难点记录
> 


```
/Users/owen/go/pkg/mod/github.com/siddontang/go-mysql@v1.1.0/replication/event.go : 103
fmt.Fprintf(w, "=== %s ===\n", EventType(h.EventType))

EventType []byte是类型 
h.EventType 是int类型？
转为EventType类型后，其有实现String方法，  在输出变量时，根据其值自动输出不同类型

变量定义在：
/Users/owen/go/pkg/mod/github.com/siddontang/go-mysql@v1.1.0/replication/const.go
```


# 关于orm golang model 结合proto的说明
> https://developpaper.com/question/how-to-use-go-language-grpc-in-combination-with-gorm/
基本上是不太能结合的。  难点有：
 - go model struct 转proto的工具不多，star最多一个已经放弃维护。 还没深入model和proto不好结合的点在哪儿，所以暂不知放弃原因。
 -  proto服务使用的机构体是根据proto文件生成。 使用orm取出来的数据赋值给 proto.go struct.(因为字段是有序的？类型是确定的？)

## protobuf time格式
> https://stackoverflow.com/questions/52802521/how-can-i-get-time-time-in-golang-protobuf-v3-struct
> https://godoc.org/github.com/golang/protobuf/ptypes#Timestamp