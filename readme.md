# 项目
业务数据模型管理系统

地址：> https://gitee.com/youwen21/business_data_model-go


## 安装
### 下载地址
> https://gitee.com/youwen21/business_data_model-go/releases

linux, mac, windows分别不同的执行文件。
release_linux-20210101.zip, release_mac-20210101.zip, release_windows-20210101.zip

解压后目录结构如下：
```
.
├── assets // html静态资源
│ ├── AdminLTE-3.0.5
│ ├── admin
│ ├── dist
│ └── plugins
├── storage // app.pid. log文件存放位置
├── business_data_model
└── business_event.sql // 项目初始化的sql文件
```
### 导入mysql
business_event.sql 文件导入到mysql数据库中

### 配置env 数据库信息
```
DB_EVENT_HOST=127.0.0.1
DB_EVENT_USERNAME=root
DB_EVENT_PASSWORD=root
DB_EVENT_DATABASE=business_event
DB_EVENT_CHARSET=utf8
DB_EVENT_FILTER=".*\\..*"
```

### 启动web服务
```bash
./business_data_model web-start 
```
本地（APP_ENV=local时）自动打页网页, 用户名密码 默认为admin, admin
> http://127.0.0.1:8000/admin/entrance/login.html
`APP_ENV不等于local时，不自动调起浏览器`

`如何修改登录密码`
> echo -n "新密码" | openssl dgst -sha1

把生成的新密码，替换administrator表中password字段既可。

### 启动mysql-binlog消费服务。 
一个mysql实例对应一个消费者。 非一个数据库对应一个消费者。  
`注意：启动mysql-binlog消费者的mysql 链接信息不需要配置到env中， 直接命令行指定。 `

```bash
./business_data_model binlog-start -host="127.0.0.1" -username="root" -password="password"
```
#### 指定过滤数据库表
默认过滤条件为: ".*\\..*" , 监听任何库，任何表的binlog事件。
启动mysql-binlog消费者时，指定过滤条件  
```bash
#多个过滤表达用，号隔开。
#codeper库的任何表，和test.table1, test.table2都是关系的表。 不符合条件的抛弃，不写入ddd_event_stream表中。
./business_data_model binlog-start -host="127.0.0.1" -username="root" -password="password" -filter="codeper\\..*,test\\.table1,test\\.table2"
```

### 可以监听多个mysql实例， 如再增加一个监听
mysql-binlog消费者可以任意启动多个。  
```bash
nohup ./business_data_model binlog-start -host="127.0.0.1" -username="root" -password="password" 2&>1 >out.log &
```

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




