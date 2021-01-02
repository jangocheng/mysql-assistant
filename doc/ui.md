## 业务数据模型管理系统元素说明
 - 事件更表（也就是数据模型）： 组成系统的重要动作， 如登录，注册，支付，支付回调等
 - binlog数据留：所有监听的库.表的插入，更新，删除操作记录。
 - 事件对比： 对多事件包含的库.表是否相同， 库表的操作字段值是否想同，如status字段
 - ER图： Entity Relationship Diagram, 数据表之间的关系，用哪些字段做的关联

## 业务数据模型管理系统使用说明
### 生成的data_stream数据流如下,可标识出同事务插入，列新，删除的数据。支持库，表，事务的搜索。
![事件数据流](/doc/images/event_stream.jpg)
### 根据数据流，创建事件
![创建事件](/doc/images/event_create.jpg)
### 查看事件或数据流内容详情
![事件和数据内容详情](/doc/images/data_model_effect.jpg)
### 数据流或事件对比创建好的事件
![数据流或事件对比事件](/doc/images/data_model_diff.jpg)

## 更新数据异同对比示例
![数据对比](/doc/images/diff_update_column.jpg)
![数据对比](/doc/images/diff_update_value.jpg)

## 根据事件ID或者勾选的事件流IDS自动生成ER图
![自动生成ER图](/doc/images/create_er.jpg)