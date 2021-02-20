# 安装项目数据模型管理系统

## 安装方式

### 编译安装

### 可执行文件包下载安装
###  下载地址
> https://gitee.com/youwen21/business_data_model-go/releases


## 目录结构重点元素：
```
.
├── assets // html静态资源，无需nginx提供完整服务
│ ├── AdminLTE-3.0.5 // 管理后台框架
│ ├── admin // 管理后台页面
│ ├── dist // 项目js，image等
│ └── plugins // js 外部包
├── storage // app.pid， log日志文件存放位置
├── xxx.exe 可执行文件
└── business_event.sql // 项目初始化的sql文件
```

## 配置环境变量

### 项目数据库配置
`注意： 被消费mysql-binlog的主库配置不需要写到env文件中。 主库信息由启动命令参数指定`
```
DB_EVENT_HOST=127.0.0.1
DB_EVENT_USERNAME=root
DB_EVENT_PASSWORD=root
DB_EVENT_DATABASE=business_event
DB_EVENT_CHARSET=utf8
DB_EVENT_FILTER=".*\\..*" //指定关心的库包，支持正则，配置方式可参考canal filter配置文式
```
### 项目功能配置
```

```


## 导入mysql
business_event.sql 文件导入到mysql数据库中

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

