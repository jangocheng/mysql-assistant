package admin

import (
	"encoding/json"
	"fmt"
	"owen2020/app/apputil"
	"owen2020/app/models"
	"owen2020/app/models/dao"
	"owen2020/app/reqt"
	"owen2020/app/resp"
	"owen2020/app/resp/out"
	"owen2020/conn"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDddEventList(c *gin.Context) {
	pageParams := &reqt.PageParam{}
	c.ShouldBindQuery(&pageParams)

	eventList := []models.DddEvent{}
	db := conn.GetEventGorm()
	query := db.Table("ddd_event").Where("is_deleted = 0")

	var total int64
	if err := query.Count(&total).Error; err != nil {
		out.NewError(700, err.Error()).JSONOK(c)
		return
	}

	err := query.Order("ddd_event_id desc").Limit(pageParams.Limit()).Offset(pageParams.Offset()).Find(&eventList).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(gin.H{"total": total, "rows": eventList}).JSONOK(c)
}

func GetDddEventInfo(c *gin.Context) {
	id := c.Param("id")
	// idInt, _ := strconv.Atoi(id)

	eventInfo := models.DddEvent{}
	db := conn.GetEventGorm()
	query := db.Table("ddd_event").Where("is_deleted = 0 and ddd_event_id = ?", id)

	err := query.First(&eventInfo).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(eventInfo).JSONOK(c)
}

//AddDddEvent 添加
func AddDddEvent(c *gin.Context) {
	eventInfo := models.DddEvent{}

	err := apputil.ShouldBindOrError(c, &eventInfo)
	if err != nil {
		return
	}

	db := conn.GetEventGorm()
	err = db.Table("ddd_event").Create(&eventInfo).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(eventInfo).JSONOK(c)
}

//EditDddEvent 编辑
func EditDddEvent(c *gin.Context) {
	id := c.Param("id")

	eventInfo := models.DddEvent{}
	err := apputil.ShouldBindOrError(c, &eventInfo)
	if err != nil {
		return
	}
	eventInfo.DddEventId, _ = strconv.Atoi(id)

	// config := &gorm.Session{DryRun: true}
	config := &gorm.Session{}
	db := conn.GetEventGorm()
	session := db.Session(config)
	stmt := session.Table("ddd_event").Select("*").Where("ddd_event_id = ?", id).UpdateColumns(eventInfo).Statement
	//stmt := session.Table("ddd_event").Where("ddd_event_id = ?", id).Save(eventInfo).Statement

	fmt.Println("sql is :", stmt.SQL.String())
	err = session.Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess(eventInfo).JSONOK(c)
}

func DeleteDddEvent(c *gin.Context) {
	id := c.Param("id")
	// idInt, _ := strconv.Atoi(id)

	// eventInfo := models.DddEvent{}
	db := conn.GetEventGorm()
	query := db.Table("ddd_event").Where("ddd_event_id = ?", id)

	err := query.Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	if err != nil {
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}

	out.NewSuccess("").JSONOK(c)
}

type sortStreams []models.DddEventStream

func (s sortStreams) Len() int { return len(s) }

func (s sortStreams) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s sortStreams) Less(i, j int) bool {
	if s[i].DbName > s[j].DbName {
		return true
	}

	if s[i].TableName > s[j].TableName {
		return true
	}

	if s[i].EventType > s[j].EventType {
		return true
	}

	return false
}

// @see https://stackoverflow.com/questions/36854408/how-to-append-to-a-slice-pointer-receiver
// @see https://www.pauladamsmith.com/blog/2016/07/go-modify-slice-iteration.html
// extractAndRemove
func (s *sortStreams) extractAndRemove(db string, table string) models.DddEventStream {
	for i, v := range *s {
		if v.DbName == db && v.TableName == table {
			// 从s中移出匹配的元素
			*s = append((*s)[:i], (*s)[i+1:]...)
			return v
		}
	}
	return s.getDefault()
}

func (s sortStreams) getDefault() models.DddEventStream {
	stream := &models.DddEventStream{}
	stream.EventType = -100
	return *stream
}

func (s sortStreams) process() {
	// 处理Columns，UpdateColumns
	// 处理update_value
	for i, v := range s {
		if v.Columns != "" {
			v.Columns = strings.ReplaceAll(v.Columns, ",", "\n")
		}
		if v.UpdateColumns != "" {
			v.UpdateColumns = strings.ReplaceAll(v.UpdateColumns, ",", "\n")
		}
		v.UpdateValue = processValues(v.UpdateValue)

		s[i] = v
	}
}

func processValues(jsonStr string) string {
	jmap := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &jmap)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	str := ""
	for k, v := range jmap {
		str += k + ":" + v + "\n"
	}
	return str
}

func EventDiff(c *gin.Context) {
	searchParams := reqt.EventDiffSearchParams(c)

	streamListA, _ := dao.GetStreamListByEventId(searchParams.EventA)
	var streamListB []models.DddEventStream

	if searchParams.EventB == "" {
		streamListB, _ = dao.GetStreamListByIds(strings.Split(searchParams.StreamIds, ","))
	} else {
		streamListB, _ = dao.GetStreamListByEventId(searchParams.EventB)
	}

	streamA := sortStreams(streamListA)
	streamB := sortStreams(streamListB)
	sort.Sort(streamA)
	sort.Sort(streamB)

	streamA.process()
	streamB.process()

	result := combineList(streamA, streamB)

	// out.NewSuccess(map[string]interface{}{"a": streamA, "b": streamB, "diff": result}).JSONOK(c)
	out.NewSuccess(result).JSONOK(c)
	// out.NewSuccess(result).JSONOK(c)
}

func combineList(a sortStreams, b sortStreams) []resp.DiffRow {
	var diffList []resp.DiffRow

	for _, v := range a {
		// diffEntityRow := &diffEntity{A: v}
		diffEntityRow := &resp.DiffRow{}
		diffEntityRow.A = resp.GenDiffEntity(v)
		diffEntityRow.B = resp.GenDiffEntity(b.extractAndRemove(v.DbName, v.TableName))
		diffList = append(diffList, *diffEntityRow)
	}

	for _, v := range b {
		diffEntityRow := &resp.DiffRow{A: resp.GenDiffEntity(b.getDefault()), B: resp.GenDiffEntity(v)}
		diffList = append(diffList, *diffEntityRow)
	}

	return diffList
}
