package routes

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"owen2020/app/http/middleware"
)

//StartGin 启动服务器，监听端口, 配置路由
func StartGin() {
	router := SetUpRouter()
	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	// endless.DefaultReadTimeOut = readTimeout
	// endless.DefaultWriteTimeOut = writeTimeout
	// endless.DefaultMaxHeaderBytes = maxHeaderBytes
	// go func() {
	// 	router.Run(":80")
	// }()
	//server := endless.NewServer(os.Getenv("APP_PORT"), router)
	//server.BeforeBegin = func(add string) {
	//	pid := syscall.Getpid()
	//	log.Printf("Actual pid is %d", pid)
	//	WritePidToFile(pid)
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}

	router.Run(os.Getenv("APP_PORT"))
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// 静态文件
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")   // 单文件
	router.Static("/dist", "./assets/dist")                     // 目录下的文件
	router.Static("/plugins", "./assets/plugins")               // 目录下的文件
	router.Static("/AdminLTE-3.0.5", "./assets/AdminLTE-3.0.5") // 目录下的文件
	router.Static("/admin", "./assets/admin")                   // 目录下的文件
	//router.StaticFS("/more_static", http.Dir("my_file_system"))  // 目录正的文件，定制file.System服务

	// admin 路由
	adminRoute(router)

	// 未匹配到任何路由
	router.NoRoute(func(c *gin.Context) {
		// router.HandleContext()
		c.AbortWithStatus(http.StatusNotFound)

	})

	return router
}

func WritePidToFile(pid int) {
	pidFile := getAppFile()

	f, err := os.OpenFile(pidFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666) //打开文件, 不存在就创建
	if err != nil {
		log.Fatal("打开文件失败：", err.Error())
	}
	defer f.Close()

	pidStr := strconv.Itoa(pid)
	ret, err := io.WriteString(f, pidStr)
	if err != nil {
		log.Fatal("写入文件失败：", err.Error())
	}
	log.Println("app.pid 写入结果：", pidStr, ",长度", ret)
}

func RmPidFile() {
	pidFile := getAppFile()
	os.Remove(pidFile)
}

func getAppFile() string {
	pwd, _ := os.Getwd()
	pidFile := pwd + "/storage/app.pid"

	return pidFile
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
