package activity

//字符串转换 https://blog.csdn.net/hanyajun0123/article/details/92445437
//go语言中获取变量类型的三种方法 https://www.cnblogs.com/-wenli/p/11436810.html
// validator自定义错误提示语
import (

	// "net/http/httputil"
	"fmt"
	"owen2020/app/apputil"
"owen2020/app/resp/out"
	"owen2020/app/apputil/applog"
	"owen2020/app/models"
	"owen2020/app/resp"
	"owen2020/conn"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	_ "net/http/httputil"
)

func GetList(c *gin.Context) {
	var activityList []models.Activity
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	xorm := conn.GetXormWithLog()
	session := xorm.Table("activity")

	total, _ := session.Count()
	err := session.Limit(pageSize, (page-1)*pageSize).OrderBy("activity_id desc").Find(&activityList)
	// result, err := xorm.QueryString("select * from activity")
	if nil != err {
		out.NewError(600, "哈哈").JSONOK(c)
		return
	}

	respList := resp.GenActivityList(activityList)
	out.NewSuccess(gin.H{"total": total, "rows": respList}).JSONOK(c)
}

func GetInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var activityInfo models.Activity

	fmt.Println(id)

	engine := conn.GetXormWithLog()
	ret, err := engine.Where("activity_id = ?", id).Get(&activityInfo)
	if nil != err {
		out.NewError(600, err.Error()).JSONOK(c)
		return
	}
	if !ret {
		out.NewError(600, "未找到数据").JSONOK(c)
		return
	}

	respInfo := resp.GenInfoByActInfo(activityInfo)
	out.NewSuccess(respInfo).JSONOK(c)
}

//Insert 创建活动
// https://github.com/gin-gonic/gin/issues/430
// https://blog.depa.do/post/gin-validation-errors-handling#toc_3
//
func Insert(c *gin.Context) {
	// activityInfo := new(models.Activity)
	var activityInfo *models.Activity = &models.Activity{}

	// activityInfo, err := request.ShouldBindActivityInfo(c)
	err := apputil.ShouldBindOrError(c, &activityInfo)
	// 由request层输出验证错误类信息
	if nil != err {
		return
	}
	memberID, has := c.Get("member_id")
	if has {
		activityInfo.MemberId, _ = strconv.Atoi(memberID.(string))
	}

	applog.Logger.WithFields(logrus.Fields{"activitiInfo": activityInfo}).Info("哈哈")
	// 插入数据
	engine := conn.GetXormWithLog()
	ret, err := engine.InsertOne(activityInfo)
	if nil != err {
		out.NewError(600, err.Error()).JSONOK(c)
		return
	}
	if 0 == ret {
		out.NewError(600, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(activityInfo).JSONOK(c)
}

//Update 更新活动
func Update(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	// activityInfo := new(models.Activity)
	var activityInfo *models.Activity = &models.Activity{}

	err := apputil.ShouldBindOrError(c, &activityInfo)
	// 由request层输出验证错误类信息
	if nil != err {
		return
	}
	memberID, _ := c.Get("member_id")
	memberIDStr, _ := strconv.Atoi(memberID.(string))
	// if has {
	// 	activityInfo.CreatedBy, _ = strconv.Atoi(memberID.(string))
	// }

	// 更新数据
	engine := conn.GetXormWithLog()
	ret, err := engine.Update(&activityInfo, &models.Activity{ActivityId: idInt, MemberId: memberIDStr})
	if nil != err {
		out.NewError(600, err.Error()).JSONOK(c)
		return
	}
	if 0 == ret {
		out.NewError(601, "更新数据为0").JSONOK(c)
		return
	}

	out.NewSuccess(activityInfo).JSONOK(c)
}

//Delete 删除活动
func Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	memberID, _ := c.Get("member_id")
	memberIDStr, _ := strconv.Atoi(memberID.(string))

	// 更新数据
	engine := conn.GetXormWithLog()

	// activityInfo := &models.Activity{ActivityId: idInt, IsDeleted: 1, DeletedAt: }
	// condition := &models.Activity{CreatedBy: memberIDStr}
	// fmt.Printf("%+v", activityInfo)
	// fmt.Printf("%+v", condition)
	// ret, err := engine.Id(idInt).Update(&activityInfo)

	sql := "update `activity` set is_deleted = ?, deleted_at=?  where activity_id = ? and created_by = ? "
	ret, err := engine.Exec(sql, 1, time.Now().Format("2006-01-02 15:04:05"), idInt, memberIDStr)

	if nil != err {
		out.NewError(600, err.Error()).JSONOK(c)
		return
	}

	if aff, _ := ret.RowsAffected(); aff == 0 {
		out.NewError(601, "无记录更新").JSONOK(c)
		return
	}

	out.NewSuccess("").JSONOK(c)
}
