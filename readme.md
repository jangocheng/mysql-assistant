# mysql助手管理系统

功能一：统计mysql每日增删改次数
功能二：检查字段变更是否合规
功能三：mysql-binlog数据流解析记录，组合生成接口文档

### 功能一：数据统计

- 设计统计规则 （库+表[增，改，删]或 库+表+字段[改]）
- 按天为单位，统计每天相应规则的增，改，删的总数

[数据变更统计说明](/doc/statistics.md)

### 功能二：状态管理，状态校验

- 定义状态节点和状态节点关系， 以图型方式直观展示节点关系
- 校验状态数据变更是否合规正确

[状态管理和变更不是合规校验](/doc/status_flow.md)

### 功能三：mysql-binlog数据流解析， 接口文档生成

- 定义接口文档
- 展示接口影响的库，表，字段
- 自动画数据模型ER图

[mysql-binlog解析说明和生成api接口](/doc/data_model.md)

## 安装使用

[安装使用说明文档](/doc/install.md)

## Docker 快速开始

[docker](/doc/docker_quick_start.md)

## 开发调试与技术实现原理

[开发调试与技术实现原理说明文档](/doc/advance.md)

## 系统依赖

目前只有态流定义和图形化展功能无任何依赖。  
其他功能都依赖MySQL数据库。


## k8s布置

deploy目录中已定义k8s deployment。  
如果使用本地minikube开发， 可参考doc目下 k8s_command和minikube_command，实现本地访问admin管理后台





