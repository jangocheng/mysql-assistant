package reqt

import "github.com/gin-gonic/gin"

//PageParam 分页控制
type PageParam struct {
	PageNumber int `json:"pageNumber" form:"pageNumber,default=1"`
	PageSize   int `json:"pageSize" form:"pageSize,default=10"`
}

//Limit 获取每页记录数
func (p *PageParam) Limit() int {
	return p.PageSize
}

//Offset 获取当前页开始ID
func (p *PageParam) Offset() int {
	if p.PageNumber == 0 {
		return 0
	}

	return (p.PageNumber - 1) * p.PageSize
}

//PageParams 分页参数获取
func PageParams(c *gin.Context) PageParam {
	page := PageParam{}
	// page.pageSize, _ = strconv.Atoi(c.DefaultQuery("page", "10"))
	c.ShouldBindQuery(&page)

	return page
}
