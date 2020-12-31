package command

import (
	"context"
	"owen2020/cmd/command/handle_binlog"
	"owen2020/cmd/command/mysqlutil"

	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/urfave/cli/v2"

	_ "github.com/go-sql-driver/mysql"
)

// mysql binlog 消费错误场景
// https://www.jianshu.com/p/ec4e626ae0b0

// MySQL二进制日志分析 - TABLE_MAP_EVENT
// https://www.cnblogs.com/little-star-2015/p/11736822.html

// 解析MySQL binlog --（4）TABLE_MAP_EVENT
// https://blog.51cto.com/yanzongshuai/2090758

// 解析MySQL binlog --（1）大致结构及event type
//https://blog.51cto.com/yanzongshuai/2085203

// https://blog.csdn.net/whatday/article/details/107918399
// golang string int int32 int64 float32 float64 time 互相转换

// Convert binary value as string to uint32 in Golang
// https://stackoverflow.com/questions/54814382/convert-binary-value-as-string-to-uint32-in-golang/54814575

//StartBinlogClient 消费mysql binlog
func StartBinlogClient(c *cli.Context) error {
	host := c.String("host")
	port := c.Int("port")
	serverID := c.Int("server_id")
	username := c.String("username")
	password := c.String("password")
	filter := c.String("filter")
	// apputil.PrettyPrint(c.Int("port"))
	// apputil.PrettyPrint(c.Int("server_id"))
	// apputil.PrettyPrint(c.String("username"))
	// apputil.PrettyPrint(c.String("password"))
	// apputil.PrettyPrint(c.String("filter"))

	cfg := replication.BinlogSyncerConfig{
		ServerID: uint32(serverID),
		Flavor:   "mysql",
		Host:     host,
		Port:     uint16(port),
		User:     username,
		Password: password,
	}

	// 定义过滤
	handle_binlog.Filter = filter
	handle_binlog.Cfg = cfg

	// 取出所有库表的字段ID对应的字段名
	// 初始化所有库.表的字段ID对应字段名
	handle_binlog.InitDBTables()

	// 获取binlog file position
	masterPosition := mysqlutil.GetMysqlPosition(cfg)

	// apputil.PrettyPrint(handle_binlog.DBTables)
	syncer := replication.NewBinlogSyncer(cfg)
	// pos, _ := strconv.ParseUint(masterPosition["Position"], 10, 32)
	u32 := uint32(masterPosition["Position"].(uint64))
	streamer, _ := syncer.StartSync(mysql.Position{masterPosition["File"].(string), u32})
	for {
		ev, _ := streamer.GetEvent(context.Background())
		// Dump event
		//ev.Dump(os.Stdout)
		handle_binlog.HandleEvent(ev)
	}
	// or you can start a gtid replication like
	// streamer, _ := syncer.StartSyncGTID(gtidSet)
	// the mysql GTID set likes this "de278ad0-2106-11e4-9f8e-6edd0ca20947:1-2"
	// the mariadb GTID set likes this "0-1-100"

	// f

	// or we can use a timeout context
	// for {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// 	ev, err := s.GetEvent(ctx)
	// 	cancel()

	// 	if err == context.DeadlineExceeded {
	// 		// meet timeout
	// 		continue
	// 	}

	// 	ev.Dump(os.Stdout)
	// }
	return nil
}
