package common

import (
	"owen2020/app/models"
	"owen2020/app/models/dao"
)

var (
	StateClasses         = make(map[string]int)
	StateClassDirections = make(map[int][]models.StateDirection)
)

var (
	StatisticsRules map[string]int = make(map[string]int)
)

func InitStatisticsRules() {
	list, _ := dao.GetStatisticsRuleList()

	for _, v := range list {
		key := GetKey(v.DbName, v.TableName, v.FieldName)
		value := v.StatisticsRuleId
		StatisticsRules[key] = value
	}
}

func InitState() {
	InitStateClass()
	InitStateDirection()
}

func InitStateClass() {
	//stateList := dao.GetStateList();
	list, _ := dao.GetStateClassList()

	for _, v := range list {
		key := GetKey(v.DbName, v.TableName, v.FieldName)
		value := v.StateClassId
		StateClasses[key] = value
	}
}

func InitStateDirection() {
	list, _ := dao.GetAllStateDirectionList()

	for _, v := range list {
		key := v.StateClassId
		SetDirection(key, v)
	}
}

func SetDirection(classId int, row models.StateDirection) {
	_, ok := StateClassDirections[classId]
	if !ok {
		StateClassDirections[classId] = []models.StateDirection{}
	}

	StateClassDirections[classId] = append(StateClassDirections[classId], row)
}
