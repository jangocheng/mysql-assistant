package routes

import (
	"owen2020/app/http/controllers/admin"
	"owen2020/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func adminRoute(router *gin.Engine) {
	adminOpenAPI := router.Group("/admin-api/v1")
	{
		adminOpenAPI.POST("/token", admin.CreateToken)
	}

	cacheAPI := router.Group("/admin-api/v1")
	cacheAPI.Use(middleware.AdminToken)
	cacheAPI.Use(middleware.BrowserCacheMiddleware)
	{
		// 菜单树
		cacheAPI.GET("/sidebar", admin.GetSideBarTree)
	}

	adminAPI := router.Group("/admin-api/v1")
	adminAPI.Use(middleware.AdminToken)
	// 框架菜单相关
	{
		// 菜单列表
		adminAPI.GET("/menu", admin.GetMenuList)
	}

	// 数据模型相关
	{
		// 事件列表 完成
		adminAPI.GET("/ddd_event", admin.GetDddEventList)
		// 事件详情  完成
		adminAPI.GET("/ddd_event/:id", admin.GetDddEventInfo)
		// 事件删除  完成
		adminAPI.DELETE("/ddd_event/:id", admin.DeleteDddEvent)

		// 事件添加
		adminAPI.POST("ddd_event", admin.AddDddEvent)
		// 事件编辑
		adminAPI.POST("ddd_event/:id", admin.EditDddEvent)

		// 事件对比 diff
		adminAPI.GET("/event_diff", admin.EventDiff)
		adminAPI.GET("/event", admin.EventDiff)

		// adminAPI.GET("/event/effect", admin.EventEffect)

		// 数据流列表 完成
		adminAPI.GET("/ddd_event_stream", admin.GetDddEventStreamList)
		// 数据流删除 完成
		adminAPI.DELETE("/ddd_event_stream/:id", admin.DeleteDddEventStream)

		// er图 实体关系图 完成
		adminAPI.GET("/ddd_er", admin.DddEr)
	}

	// 业务状态相关
	{
		adminAPI.GET("/state_class", admin.GetStateClassList)
		adminAPI.GET("/state_class/:id", admin.GetStateClassInfo)
		adminAPI.POST("/state_class/:id", admin.EditStateClass)
		adminAPI.DELETE("/state_class/:id", admin.DeleteStatClass)

		adminAPI.GET("/state", admin.GetStateList)
		adminAPI.GET("/state/:id", admin.GetStateInfo)
		adminAPI.POST("/state/:id", admin.EditState)
		adminAPI.DELETE("/state/:id", admin.DeleteStat)

		adminAPI.GET("/state_direction", admin.GetStateDirectionList)
		adminAPI.DELETE("/state_direction/:id", admin.DeleteStatDirection)

		adminAPI.GET("/state_abnormal", admin.GetStateAbnormalList)
		adminAPI.DELETE("/state_abnormal/:id", admin.DeleteStatAbnormal)

		adminAPI.GET("/state_graph/:id", admin.StateGraph)
	}
	// 统计相关
	{
		adminAPI.GET("/statistics_rule", admin.GetStatisticRuleList)
		adminAPI.DELETE("/statistics_rule/:id", admin.DeleteStatisticRule)

		adminAPI.GET("/statistics_day", admin.GetStatisticDayList)
		//adminAPI.DELETE("/statistics_day/:id", admin.DeleteStatisticDay)
	}

}
