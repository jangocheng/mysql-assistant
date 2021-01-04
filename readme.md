# 项目

业务数据模型管理系统是一款技术管理工具，辅助技术人员管理业务数据结构。 管理业务事件和文档。

地址：> https://gitee.com/youwen21/business_data_model-go

## 解决什么问题

管理业务相关的用户事件，和事件迭代多个版本对应的数据结构。 弥补API接口文档，详细设计文档和ER图对系统描述不足。 对业务数据库结构的管理与共识，减少bug。

## 项目依赖项

- mysql数据库

## 实现说明

项目由两个重压模块组成

- mysql-binlog消费模块， 模拟mysql-slave消费mysql-binlog，生成insert,update,delete操作数据流。
- web-ui模块,启动web-server, 提供web后台管理界面。根据数据流生成数据模型，对比模型异同，输出diff对比。

## 安装

### 下载地址

> https://gitee.com/youwen21/business_data_model-go/releases

linux, mac, windows平台执行文件已编译完成。  
release_linux-20210101.zip,  
release_mac-20210101.zip,  
release_windows-20210101.zip

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
`APP_ENV不等于local时，不自动调起浏览器`
> http://127.0.0.1:8000/admin/entrance/login.html

`修改登录密码： 把生成的新密码，替换administrator表中password字段既可。`
> echo -n "新密码" | openssl dgst -sha1

### 启动mysql-binlog消费服务。

一个mysql实例对应一个消费者。 非一个数据库对应一个消费者。  
`注意：启动mysql-binlog消费者的mysql 链接信息不需要配置到env中， 直接命令行指定。 `

```bash
./business_data_model binlog-start -host="127.0.0.1" -username="root" -password="password"
```

#### 启动mysql-binlog消费服务，指定过滤数据库表

默认过滤条件为: ".*\\..*" , 监听任何库，任何表的binlog事件。 启动mysql-binlog消费者时，指定过滤条件

```bash
#多个过滤表达用，号隔开。
#codeper库的任何表，和test.table1, test.table2都是关系的表。 不符合条件的抛弃，不写入ddd_event_stream表中。
./business_data_model binlog-start -host="127.0.0.1" -username="root" -password="password" -filter="codeper\\..*,test\\.table1,test\\.table2"
```

## 监听多个mysql实例

mysql-binlog消费者可以任意启动多个。

```bash
nohup ./business_data_model binlog-start -host="127.0.0.1" -username="root" -password="password" 2&>1 >out.log &
```

## 更多文档
[业务数据模型管理系统使用说明](/doc/ui.md)





