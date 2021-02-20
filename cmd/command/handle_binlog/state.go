package handle_binlog

import (
	"errors"
	"owen2020/app/models"
	"owen2020/app/models/dao"
)

var (
	StateClasses = make(map[string]int)
	StateClassDirections = make(map[int][]models.StateDirection)
)

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

func CheckStatDirection(classId int, from string, to string) (bool, error) {
	list, ok := StateClassDirections[classId]
	if !ok {
		return false, errors.New("state class direction not exist")
	}

	for _, v := range list {
		if v.StateFrom == from && v.StateTo == to {
			return true, nil
		}
	}

	return false, errors.New("direction not defined")
}

func CheckDirection(dbName string, tableName string, fieldName string, from string, to string) (bool, error) {
	classId, err := GetStatClassId(dbName, tableName, fieldName)
	if nil != err {
		return false, errors.New("state class not defined")
	}

	list, ok := StateClassDirections[classId]
	if !ok {
		return false, errors.New("state class directions not defined")
	}

	for _, v := range list {
		if v.StateFrom == from && v.StateTo == to {
			return true, nil
		}
	}

	return false, errors.New("direction not exist")
}

func GetStatClassId(dbName string, tableName string, fieldName string) (int, error) {
	key := GetKey(dbName, tableName, fieldName)
	classId, ok := StateClasses[key]
	if !ok {
		return 0, errors.New("state class not defined")
	}

	return classId, nil
}

func GetKey(dbName string, tableName string, fieldName string) string {
	return dbName + "_" + tableName + "_" + fieldName
}
