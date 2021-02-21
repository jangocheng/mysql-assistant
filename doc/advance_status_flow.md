# 状态管理 - 开发说明

## 模块

### web-ui

### binlog-consumer
env配置项ENABLE_DATA_STATISTICS决定是否开启库表统计功能

```
# 数据统计
## 统计数据变更次数是否开启
ENABLE_DATA_STATISTICS=yes
## 统计数据累计次数， 满足立即写入到表statistics_day
DATA_STATISTICS_EVENT_TIMES=2
## 统计数据刷新频率设置 单位秒, 满足立即写入到表statistics_day
DATA_STATISTICS_FLUSH_DURATION=500
```

## 数据库表

- stat_class
- stat
- stat_direction
- stat_abnormal

## 依赖项

- mysql-binlog
- d3-graphviz

## 开发入口

- cmd/cmd.go

## 调用路径

cmd/cmd.go -> StartBinlogClient -> handle_binlog.HandleEvent(ev) -> handleUpdateEventV1(e) -> updateRoutineStatRule(ev)

