
## 实现说明
项目由两个重压模块组成
- mysql-binlog消费模块， 模拟mysql-slave消费mysql-binlog，生成insert,update,delete操作数据流。
- web-ui模块,启动web-server, 提供web后台管理界面。根据数据流生成数据模型，对比模型异同，输出diff对比。
