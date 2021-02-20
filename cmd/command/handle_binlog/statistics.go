package handle_binlog

import (
	"errors"
	"fmt"
	"os"
	"owen2020/app/models"
	"owen2020/app/models/dao"
	"owen2020/conn"
	"strconv"
	"sync"
	"time"
)

type MySyncMap struct {
	sync.Map
}

var StatisticsRules map[string]int = make(map[string]int)
var StatisticsDayData MySyncMap

var statisticsEventTimes int
var statisticsLastUpdateTime time.Time

const (
	INSERT = 1 //iota
	UPDATE = 2 //
	DELETE = 3
)

func (m MySyncMap) Print(k interface{}) {
	value, ok := m.Load(k)
	fmt.Println(value, ok)
}

func InitStatisticsRules() {
	list, _ := dao.GetStatisticsRuleList()

	for _, v := range list {
		key := GetKey(v.DbName, v.TableName, v.FieldName)
		value := v.StatisticsRuleId
		StatisticsRules[key] = value
	}
}

func GetRuleId(dbName string, tableName string, fieldName string) (int, error) {
	key := GetKey(dbName, tableName, fieldName)
	ruleId, ok := StatisticsRules[key]
	if !ok {
		return 0, errors.New("rule not exist")
	}

	return ruleId, nil
}

func StatisticsIncrease(dbName string, tableName string, fieldName string, eventType int, times int) {
	ruleID, _ := GetRuleId(dbName, tableName, fieldName)
	if ruleID == 0 {
		return
	}

	now := time.Now()
	dayKey := GetDayKey(dbName, tableName, fieldName, now)
	var dayData models.StatisticsDay

	data, ok := StatisticsDayData.Load(dayKey)
	if !ok {
		dayData = solveStatisticsDayData(dbName, tableName, fieldName, now, ruleID)
	} else {
		dayData = data.(models.StatisticsDay)
	}

	switch eventType {
	case INSERT:
		dayData.InsertTimes += times
	case UPDATE:
		dayData.UpdateTimes += times
	case DELETE:
		dayData.DeleteTimes += times
	default:
		fmt.Println("event type not support", eventType, times)
	}

	StatisticsDayData.Store(dayKey, dayData)

	statisticsEventTimes += times
	if statisticsNeedUpdate() {
		statisticsStore()
	}
}

func statisticsNeedUpdate() bool {
	eventThreshold := 500
	envTimes := os.Getenv("DATA_STATISTICS_EVENT_TIMES")
	if envTimes != "" {
		eventThreshold, _ = strconv.Atoi(envTimes)
	}
	if statisticsEventTimes > eventThreshold {
		return true
	}

	if statisticsLastUpdateTime.IsZero() {
		return false
	}

	var durationThreshold int64 = 300
	envDuration := os.Getenv("DATA_STATISTICS_FLUSH_DURATION")
	if envDuration != "" {
		durationThreshold, _ = strconv.ParseInt(envDuration, 10, 64)
	}

	if time.Now().Unix()-statisticsLastUpdateTime.Unix() > durationThreshold {
		return true
	}

	return false
}

func statisticsStore() {
	gorm := conn.GetEventGorm()
	f := func(k, v interface{}) bool {
		//这个函数的入参、出参的类型都已经固定，不能修改
		//可以在函数体内编写自己的代码，调用map中的k,v
		dayData := v.(models.StatisticsDay)
		if dayData.StatisticsDayId == 0 {
			gorm.Table("statistics_day").Create(&dayData)
			StatisticsDayData.Store(k, dayData)
		} else {
			gorm.Table("statistics_day").Save(dayData)
		}

		// 从map中删除过期数据
		if time.Now().Unix()-time.Time(dayData.StatisticsDay).Unix() > 86400 {
			StatisticsDayData.Delete(k)
		}

		return true
	}
	StatisticsDayData.Range(f)

	// 重置记数器
	statisticsEventTimes = 0
	statisticsLastUpdateTime = time.Now()
}

func GetDayKey(dbName string, tableName string, fieldName string, now time.Time) string {
	date := now.Format("20060102")
	key := GetKey(dbName, tableName, fieldName)
	return key + "_" + date
}

func solveStatisticsDayData(dbName string, tableName string, fieldName string, now time.Time, ruleId int) models.StatisticsDay {
	info, err := getFromDb(dbName, tableName, fieldName, now)
	if nil == err {
		return info
	}

	return models.StatisticsDay{
		StatisticsRuleId: ruleId,
		StatisticsDay:    models.Date(now),
		DbName:           dbName,
		TableName:        tableName,
		FieldName:        fieldName,
	}
}

func getFromDb(dbName string, tableName string, fieldName string, now time.Time) (models.StatisticsDay, error) {
	date := now.Format("2006-01-02")
	info := models.StatisticsDay{}
	gorm := conn.GetEventGorm()
	sqlWhere := "db_name = ? and table_name = ? and field_name = ? and statistics_day = ?"
	err := gorm.Table("statistics_day").Where(sqlWhere, dbName, tableName, fieldName, date).First(&info).Error
	return info, err
}
