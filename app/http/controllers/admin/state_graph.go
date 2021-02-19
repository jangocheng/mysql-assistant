package admin

import (
	"fmt"
	"github.com/emicklei/dot"
	"github.com/gin-gonic/gin"
	"owen2020/app/models"
	"owen2020/conn"
	"strconv"
)

func StateGraph(c *gin.Context) {
	id := c.Param("id")
	// idInt, _ := strconv.Atoi(id)
	stateClassId, _ := strconv.Atoi(id)

	nodeList, _ := getNodeList(stateClassId)
	directionData, _ := getDirectionData(stateClassId)

	var nodes = make(map[string]dot.Node)
	//for _, v := range nodeList {
	//	nodes[v.StateValue] = "(" + v.StateValue + ")" + v.StateValueDesc
	//}

	g := dot.NewGraph(dot.Directed)
	for _, v := range nodeList {
		node := g.Node("(" + v.StateValue + ")" + v.StateValueDesc)
		nodes[v.StateValue] = node
	}

	for _, value := range directionData {
		from, ok := nodes[value.StateFrom]
		to, ok2 := nodes[value.StateTo]
		if !ok || !ok2 {
			fmt.Println("state class miss direction node ", stateClassId, value.StateFrom, value.StateTo)
		}

		g.Edge(from, to)
	}

	output := g.String()
	fmt.Println(output)
	c.String(200, output)
}

func getNodeList(stateClassId int) ([]models.State, error) {
	list := []models.State{}

	db := conn.GetEventGorm()
	query := db.Table("state").Where("is_deleted = ?", 0)
	if stateClassId != 0 {
		query.Where("state_class_id", stateClassId)
	}

	err := query.Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func getDirectionData(stateClassId int) ([]models.StateDirection, error) {
	list := []models.StateDirection{}

	db := conn.GetEventGorm()
	query := db.Table("state_direction").Where("is_deleted = ?", 0)
	if stateClassId != 0 {
		query.Where("state_class_id", stateClassId)
	}

	err := query.Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
