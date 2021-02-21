# 数据变更统计 - 开发说明

## 模块

### web-ui

- 业务状态管理
    + 业务状态定义
    + 异常状态变更

### binlog-consumer

env配置项ENABLE_CHECK_STATE决定是否开启状态校验功能

```
# 状态管理
## 状态流程检查是否开启
ENABLE_CHECK_STATE=yes
```

## 数据库表

- statistics_rule
- statistics_day

## 依赖项

- mysql-binlog

## 开发文件入口

- cmd/cmd.go

## 调用路径

cmd/cmd.go -> StartBinlogClient -> handle_binlog.HandleEvent(ev) -> handleUpdateEventV1(e) -> updateRoutineStatistics(
ev)

