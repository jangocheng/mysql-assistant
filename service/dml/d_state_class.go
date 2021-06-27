package dml

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"owen2020/app/models"
	"owen2020/conn"
)

type stateClassDAO struct{}

var (
	StateClassDAO = &stateClassDAO{}
)

// 通过ID批量获取
func (s *stateClassDAO) GetMulti(idList []int) (map[int]*models.StateClass, error) {
	var stateMap = make(map[int]*models.StateClass, 0)
	var states = make([]models.StateClass,0 , 10)
	conn.InitEventGormPool()
	m := s.newEngine(false)
	query := m.Table("state_class").Where("state_class_id IN ?", idList)
	err:= query.Find(&states).Error
	if err != nil {
		logrus.Errorf("partnerDAO, error when GetMulti: %v", err)
	}
	for _, stateClass := range states {
		tmp := stateClass
		stateMap[stateClass.StateClassId] = &tmp
	}
	return stateMap, nil
}

// 获取引擎
func (s *stateClassDAO) newEngine(isMaster bool) *gorm.DB {
	if isMaster {
		return conn.GetEventGorm()
	} else {
		return conn.GetEventGorm()
	}
}