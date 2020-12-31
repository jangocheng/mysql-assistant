package dev

import (
	"database/sql/driver"
	"fmt"
	"owen2020/app/resp/out"
	"owen2020/conn"
	"time"

	"github.com/gin-gonic/gin"
)

// type myTime time.Time

// var _ json.Unmarshaler = &myTime{}

// func (t *myTime) MarshalJSON() ([]byte, error) {
// 	// 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
// 	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).Format("2006-01-02 15:04:05"))), nil
// }

// func (mt *myTime) UnmarshalJSON(bs []byte) error {
// 	var s string
// 	err := json.Unmarshal(bs, &s)
// 	if err != nil {
// 		return err
// 	}
// 	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
// 	if err != nil {
// 		return err
// 	}
// 	*mt = myTime(t)
// 	return nil
// }

// type Marshaler interface {
//     MarshalJSON() ([]byte, error)
// }

type JSONTime time.Time

//UnmarshalJSON json转struct
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	stringData := string(data)
	if len(stringData) == 0 || stringData == "null" || stringData == "0000-00-00 00:00:00" {
		*t = JSONTime(time.Time{})
		return nil
	}

	// Fractional seconds are handled implicitly by Parse.
	withNanos := "2006-01-02 15:04:05"
	time, err := time.Parse(`"`+withNanos+`"`, stringData)
	*t = JSONTime(time)

	return err
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	fmt.Println("输出时间看看是啥, 这个地方调用了Sting接口", t)

	if time.Time(t).IsZero() {
		// return []byte("\"\""), nil
		return []byte("\"0000-00-00 00:00:00\""), nil
	}

	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (dt JSONTime) String() string {
	t := time.Time(dt)
	if t.IsZero() {
		return "\"0000-00-00 00:00:00\""
	}

	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return "\"0000-00-00 00:00:00\""
	}

	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return stamp
}

//Scan gorm reciev
func (t *JSONTime) Scan(value interface{}) error {
	byt, ok := value.([]byte)
	if !ok {
		fmt.Println("value convert to string failed ", value)
		// return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	stringTime := string(byt)
	// fmt.Println("string time is :", stringTime)

	if stringTime == "0000-00-00 00:00:00" {
		*t = JSONTime(time.Time{})
		return nil
	}

	ttime, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, time.UTC)
	*t = JSONTime(ttime)
	return err
}

// Value gorm save
// Value return json value, implement driver.Valuer interface
func (t JSONTime) Value() (driver.Value, error) {
	ttime := time.Time(t)
	if ttime.IsZero() {
		return "", nil
	}

	//do your serializing here
	stamp := fmt.Sprintf("%s", ttime.Format("2006-01-02 15:04:05"))
	return stamp, nil
}

type test struct {
	Id        int      `json:"id"`
	CreatedAt JSONTime `json:"created_at"`
	UpdatedAt JSONTime `json:"updated_at"`
}

func DatetimeType(c *gin.Context) {
	enter := &test{}
	fmt.Printf("%+v", enter)
	err := c.ShouldBind(&enter)
	if err != nil {
		fmt.Println(err)
		out.NewError(800, err.Error()).JSONOK(c)
		return
	}
	fmt.Printf("%+v", enter)

	gorm := conn.GetGormWithLog()
	err = gorm.Table("user").Create(&enter).Error
	fmt.Printf("%+v", enter)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("没错误")
	}

	out.NewSuccess(enter).JSONOK(c)
}

func DatetimeType3(c *gin.Context) {

	enter := &test{}
	fmt.Printf("%+v", enter)

	gorm := conn.GetGormWithLog()
	err := gorm.Table("user").Where("id = ?", 5).First(&enter).Error
	fmt.Printf("%+v", enter)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("没错误")
	}

	out.NewSuccess(enter).JSONOK(c)
}

func DatetimeType2(c *gin.Context) {
	var jTime JSONTime
	jTime = JSONTime(time.Now())
	enter := &test{Id: 5, CreatedAt: jTime}

	fmt.Printf("%+v", enter)

	gorm := conn.GetGormWithLog()
	err := gorm.Table("user").Create(enter).Error

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("没错误")
	}

	out.NewSuccess(enter).JSONOK(c)
}
