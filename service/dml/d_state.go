package dml

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"owen2020/app/apputil"
	"owen2020/app/models"
	"owen2020/conn"
)

type stateDAO struct{}

var (
	StateDAO = &stateDAO{}
)

// 插入
//func (s *stateDAO) Insert(state *models.State) (*models.State, error) {
//
//}

// 删除
//func (s *stateDAO) Delete(id int) (int64, error) {
//
//}

// 更新
//func (s *stateDAO) Update(partner *mlcfg.Partner) (*mlcfg.Partner, error) {
//
//}

// 根据合作方ID获取合作方信息
//func (this *stateDAO) Get(id int, isMaster bool) (*mlcfg.Partner, error) {
//
//}

// 通过ID批量获取
func (s *stateDAO) GetMulti(idList []int) (map[int]*models.State, error) {
	var stateMap = make(map[int]*models.State, 0)
	var states = make([]models.State,0 , 10)
	conn.InitEventGormPool()
	apputil.PrettyPrint(stateMap)
	apputil.PrettyPrint(states)
	m := s.newEngine(false)
	apputil.PrettyPrint(m)
	query := m.Table("state").Where("state_id IN ?", idList)
	err:= query.Find(&states).Error
	if err != nil {
		logrus.Errorf("partnerDAO, error when GetMulti: %v", err)
	}
	for _, state := range states {
		stateMap[state.StateId] = &state
	}
	return stateMap, nil
}

// 获取合作方列表
//func (s *stateDAO) GetList(status int) ([]*mlcfg.Partner, error) {
//
//}

// 查询列表
//func (s *stateDAO) Query(params *cfg_define.PartnerQForm) (*cfg_define.PartnerQResult, error) {
//
//}

// 获取引擎
func (s *stateDAO) newEngine(isMaster bool) *gorm.DB {
	if isMaster {
		return conn.GetEventGorm()
	} else {
		return conn.GetEventGorm()
	}
}
