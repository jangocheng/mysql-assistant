# 项目
业务数据模型管理系统是一款技术管理工具，辅助技术人员管理业务数据结构。 管理业务事件和文档。  

地址：> https://gitee.com/youwen21/business_data_model-go

## 解决什么问题
当下用户下单事件都插入了哪些表？ 更新哪些表？ 更新的哪些字段？  
每个事件迭代过几版？    
新来的开发人员要怎么熟悉原来的系统？  想知道用户下单做了哪些事情必须要看源码吗？  
因快束发展，业务数据库中是否有太多冗余字段？  
系统BUG多数因数据不正确，或导致数据不正确，有没有办法降低错误数据机率？  
一个功能，数据库应该修改哪几个字段，取哪些字段，是否需要口传？ 是否可让新接手的人更容易了解系统？   
  
有以上疑问的任何一条， 都可以用业务数据模型管理解决。

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

[业务数据模型管理系统使用说明](/doc/ui.md)





