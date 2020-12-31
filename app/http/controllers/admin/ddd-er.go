package admin

import (
	"fmt"
	"owen2020/app/models"
	"owen2020/app/models/dao"
	"owen2020/app/reqt"
	"owen2020/app/resp/out"
	"strings"

	"github.com/gin-gonic/gin"
)

func DddEr(c *gin.Context) {
	search := reqt.ErSearchParams(c)

	var list []models.DddEventStream
	var err error
	switch search.Type {
	case "stream_ids":
		list, err = dao.GetStreamListByIds(search.SearchIdSlice())
		break
	case "event_id":
		list, err = dao.GetStreamListByEventId(search.Search)
		break
	case "transaction_id":
		list, err = dao.GetStreamListByTransactionId(search.Search)
		break
	default:
		out.NewError(800, "不支持类型"+search.Type).JSONOK(c)
		break
	}
	fmt.Println(err)

	data := processStreamListToErData(list, search.Scope, true)

	// erList := ""
	out.NewSuccess(data).JSONOK(c)
}

type EntityItem struct {
	Name  string `json:"name"`
	IsKey bool   `json:"isKey"`
}

type Entity struct {
	Key   string       `json:"key"`
	Items []EntityItem `json:"items"`
}

type linkData struct {
	key     string
	columns []string
}

type link struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

var ignoreField []string = []string{"deleted_at", "created_at", "updated_at", "is_deleted", "status", "sys_update_dc"}

func processStreamListToErData(list []models.DddEventStream, scope string, filter bool) map[string]interface{} {
	entityList := []Entity{}

	var links []linkData

	for _, value := range list {
		columns := extractColumns(&value, scope)
		entityInfo := &Entity{}
		entityInfo.Key = value.DbName + "." + value.TableName

		keyName := value.TableName + "_id"
		for _, fieldName := range columns {
			item := &EntityItem{Name: fieldName, IsKey: fieldName == keyName}
			entityInfo.Items = append(entityInfo.Items, *item)
		}

		entityList = append(entityList, *entityInfo)

		// 组装linkData数据
		linkd := &linkData{key: entityInfo.Key}
		if filter {
			linkd.columns = filterColumns(columns)
		} else {
			linkd.columns = columns
		}
		links = append(links, *linkd)
	}

	link := genRealLink(links)

	return map[string]interface{}{"nodeData": entityList, "linkData": link}
}

func array_intersect(a []string, b []string) []string {
	var newS []string
	for _, av := range a {
		for _, bv := range b {
			if av == bv {
				newS = append(newS, av)
			}
		}
	}

	return newS
}

func genRealLink(d []linkData) []link {
	var linkEntitys []link
	len := len(d)
	if len < 1 {
		return linkEntitys
	}

	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			remainColumns := array_intersect(d[i].columns, d[j].columns)
			linkRow := &link{From: d[i].key, To: d[j].key, Text: strings.Join(remainColumns, ",")}
			linkEntitys = append(linkEntitys, *linkRow)
		}
	}

	return linkEntitys
}

func extractColumns(info *models.DddEventStream, scope string) []string {
	switch scope {
	case "columns":
		return strings.Split(info.Columns, ",")
	case "update_columns":
		return strings.Split(info.UpdateColumns, ",")
	}

	return nil
}

func filterColumns(columns []string) []string {
	var remain []string
	for _, name := range columns {
		if inIgnoreField(name) == false {
			remain = append(remain, name)
		}
	}

	return remain
}

func inIgnoreField(name string) bool {
	for _, ignore := range ignoreField {
		if name == ignore {
			return true
		}
	}

	return false
}
