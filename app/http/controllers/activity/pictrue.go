package activity

import (
	"os"
	"owen2020/app/resp/out"
	"owen2020/conn"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//PictrueUpload 上传图版到ali oss
func PictrueUpload(c *gin.Context) {
	memberID, has := c.Get("member_id")
	if !has {
		out.NewError(700, "权限不允许").JSONOK(c)
		return
	}

	file, _ := c.FormFile("file")
	reader, err := file.Open()
	defer reader.Close()
	// Upload the file to specific dst.
	// c.SaveUploadedFile(file, "/Users/owen/go/src/learn-go/storage/logs/aaa.png")
	unixTime := time.Now().Unix()
	filename := "uploads/activity_" + memberID.(string) + "_" + strconv.FormatInt(unixTime, 10) + "_" + file.Filename

	bucket, err := conn.OssUseDefaultBucket()
	if err != nil {
		out.NewError(700, "初始化对象存储失败:"+err.Error()).JSONOK(c)
		return
	}

	if err = bucket.PutObject(filename, reader); err != nil {
		out.NewError(700, "上传失败:"+err.Error()).JSONOK(c)
		return
	}
	// err = bucket.PutObjectFromFile(filename, "/Users/owen/go/src/learn-go/storage/logs/aaa.png")
	// if err != nil {
	// 	out.NewError(700, "上传失败:"+err.Error()).JSONOK(c)
	// 	return
	// }
	httpHost := os.Getenv("OSS_HTTP_HOST")

	out.NewSuccess(gin.H{"file": httpHost + filename}).JSONOK(c)
}
