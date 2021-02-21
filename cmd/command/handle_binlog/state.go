package handle_binlog

import (
	"errors"
	"owen2020/cmd/command/handle_binlog/common"
)

func CheckStatDirection(classId int, from string, to string) (bool, error) {
	list, ok := common.StateClassDirections[classId]
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

//func CheckDirection(dbName string, tableName string, fieldName string, from string, to string) (bool, error) {
//	classId, err := GetStatClassId(dbName, tableName, fieldName)
//	if nil != err {
//		return false, errors.New("state class not defined")
//	}
//
//	list, ok := StateClassDirections[classId]
//	if !ok {
//		return false, errors.New("state class directions not defined")
//	}
//
//	for _, v := range list {
//		if v.StateFrom == from && v.StateTo == to {
//			return true, nil
//		}
//	}
//
//	return false, errors.New("direction not exist")
//}

func GetStatClassId(dbName string, tableName string, fieldName string) (int, error) {
	key := common.GetKey(dbName, tableName, fieldName)
	classId, ok := common.StateClasses[key]
	if !ok {
		return 0, errors.New("state class not defined")
	}

	return classId, nil
}
