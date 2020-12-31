package member

import (
	"fmt"
	"owen2020/app/apputil"
"owen2020/app/resp/out"
	"owen2020/app/apputil/foundations"
	"owen2020/app/models"
	"owen2020/conn"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//GetMemberList 获取列表
func GetMemberList(c *gin.Context) {
	return
}

//GetMemberInfo 获取个人详情
func GetMemberInfo(c *gin.Context) {
	id := c.Param("id")
	authedID, _ := c.Get("member_id")
	if authedID != id {
		fmt.Println(id, "===", authedID)
		out.NewError(700, "不允许的请求").JSONOK(c)
		return
	}
	var member models.Member
	// sqlx := conn.GetSQLx()
	// err := sqlx.Get(&member, "select * from `member` where member_id=?", id)

	engine := conn.GetXorm()
	has, err := engine.Table("member").Where("member_id = ?", id).Get(&member)
	if false == has {
		out.NewError(700, "没有相关记录").JSONOK(c)
		return
	}

	if nil != err {
		out.NewError(700, "get error: "+err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(member).JSONOK(c)
	return
}

//AddMember 用户添加
func AddMember(c *gin.Context) {
	var member models.Member
	err := apputil.ShouldBindOrError(c, &member)
	if nil != err {
		return
	}

	engine := conn.GetXormWithLog()
	member.Password = foundations.EncryptWord(member.Password, nil)
	_, err = engine.Insert(&member)
	if err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
	}

	out.NewSuccess(gin.H{"member": member, "err": err}).JSONOK(c)
	return
}

//AddMemberAndLogin 用户注册
func AddMemberAndLogin(c *gin.Context) {
	var member models.Member
	err := apputil.ShouldBindOrError(c, &member)
	if nil != err {
		return
	}

	engine := conn.GetXormWithLog()
	member.Password = foundations.EncryptWord(member.Password, nil)
	affected, err := engine.Insert(&member)
	if err != nil {
		out.NewError(700, "affected:"+err.Error()).JSONOK(c)
		return
	}
	if affected == 0 {
		out.NewError(700, "手机号已存在，不可重复注册").JSONOK(c)
		return
	}

	tokenString, _ := genMemberToken(strconv.Itoa(member.MemberId))

	// 设置cookie
	c.SetCookie("Authorization", tokenString, 1000, "/", c.Request.Host, false, true)

	member.Password = ""
	// 如果是ajax，不需要跳转，直接输出结果
	if strings.EqualFold(c.Request.Header.Get("X-Requested-With"), "XMLHttpRequest") {
		out.NewSuccess(gin.H{"token": tokenString, "info": member, "jump_url": ""}).JSONOK(c)
		return
	}
	// 登录完成跑转到jump_url
	c.Redirect(302, "")
}

//EditMember 用户编辑
func EditMember(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	var member models.Member
	c.ShouldBind(&member)

	authedID, _ := c.Get("member_id")
	if authedID != id {
		out.NewError(700, "不允许的请求"+id+"="+authedID.(string)).JSONOK(c)
		return
	}

	engine := conn.GetXormWithLog()
	// _, err := engine.ID(id).Update(&member)
	_, err := engine.Update(&member, &models.Member{MemberId: idInt})
	if nil != err {
		out.NewError(700, "更新错误:"+err.Error()).JSONOK(c)
	}

	out.NewSuccess("").JSONOK(c)
	return
}

//DeleteMember  用户删除
func DeleteMember(c *gin.Context) {

}
