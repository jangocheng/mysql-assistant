package env

// 从配置中心拉取配置 覆盖默认env
// 服务容器化需要把mysql, redis配置读取。

// 服务发现不用写到这个项目里。
// 这个项目使用者应该有页面直接修改env配置？

// 使用者可以 export环境变量， 也可以使用env,  如果单机docker 可以 ENV配置数据库信息