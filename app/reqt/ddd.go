package reqt

import (
	"strings"

	"github.com/gin-gonic/gin"
)

//DddErSearch DddEr搜索
type DddErSearch struct {
	Search string `json:"search" form:"search"`
	Type   string `json:"type" form:"type"`
	Scope  string `json:"scope" form:"scope"`
}

// 事件对比搜索条件
type EventDiffSearch struct {
	EventA    string `json:"event_a" form:"event_a"`
	EventB    string `json:"event_b" form:"event_b"`
	StreamIds string `json:"stream_ids" form:"stream_ids"`
}

func (er *DddErSearch) SearchIdSlice() []string {
	return strings.Split(er.Search, ",")
}

func ErSearchParams(c *gin.Context) *DddErSearch {
	search := &DddErSearch{}
	c.ShouldBindQuery(&search)

	if search.Type == "" {
		search.Type = "event_id"
	}

	if search.Scope == "" {
		search.Scope = "columns"
	}

	return search
}

func EventDiffSearchParams(c *gin.Context) *EventDiffSearch {
	search := &EventDiffSearch{}
	c.ShouldBindQuery(&search)

	return search
}
