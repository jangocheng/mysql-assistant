package resp

import (
	"owen2020/app/apputil"
	"owen2020/app/models"

	com_alibaba_otter_canal_protocol "github.com/withlin/canal-go/protocol"
)

var statusMap map[int]string = map[int]string{
	0: "默认",
	1: "嘿嘿",
	2: "哈哈",
}

//EventStream 输出模型异同对比
type EventStream struct {
	models.DddEventStream
	EventName string `json:"event_name"`
}

//DiffRow 一行异同对比
type DiffRow struct {
	A EventStream `json:"a"`
	B EventStream `json:"b"`
}

func GenDiffEntity(m models.DddEventStream) EventStream {
	str := com_alibaba_otter_canal_protocol.EventType(m.EventType).String()
	return EventStream{m, str}
}

//GetDiffStatusText 根据int转为文本
func GetDiffStatusText(status int) string {
	return apputil.MapedIntOrDefault(statusMap, status)
}
