package store

import (
	"os"
	"owen2020/app/models"
	"owen2020/conn"
	"strconv"
	"sync"
	"time"
)

var (
	// 逻辑中使用的某个变量
	eventStreamData           []models.DddEventStream = make([]models.DddEventStream, 0, 200)
	eventStreamLastUpdateTime time.Time
	// 与变量对应的使用互斥锁
	dataGuard sync.Mutex
)

func streamNeedUpdate() bool {
	defaultMaxRows := 200
	fRows := os.Getenv("MODEL_STREAM_FLUSH_ROWS")
	if fRows != "" {
		defaultMaxRows, _ = strconv.Atoi(fRows)
	}
	if len(eventStreamData) > defaultMaxRows {
		return true
	}

	if eventStreamLastUpdateTime.IsZero() {
		return false
	}

	var durationThreshold int64 = 300
	envDuration := os.Getenv("MODEL_STREAM_FLUSH_DURATION")
	if envDuration != "" {
		durationThreshold, _ = strconv.ParseInt(envDuration, 10, 64)
	}

	if time.Now().Unix()-eventStreamLastUpdateTime.Unix() > durationThreshold {
		return true
	}

	return false
}

func StreamAddRows(ss []models.DddEventStream) {
	streamAddToMem(ss)

	if !streamNeedUpdate() {
		return
	}

	ok, _ := streamStore()
	if ok {
		eventStreamLastUpdateTime = time.Now()
	}
}

func streamAddToMem(ss []models.DddEventStream) {
	// 锁定
	dataGuard.Lock()
	// 在函数退出时解除锁定
	defer dataGuard.Unlock()

	eventStreamData = append(eventStreamData, ss...)
}

func streamStore() (bool, error) {
	dataGuard.Lock()
	defer dataGuard.Unlock()

	gorm := conn.GetEventGorm()
	gorm.Table("ddd_event_stream").Create(eventStreamData)

	//eventStreamData = make([]models.DddEventStream, 200)
	eventStreamData = eventStreamData[0:0]

	return true, nil
}
