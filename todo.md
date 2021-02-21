# TODO

### k8s发布服务
文档 https://www.kubernetes.org.cn/docs   
- 发布多个web服务 提供网页端。  Services+ReplicaSets
- 发布一个消费mysql-binlog服务  使用ReplicaSets启动一个Pod
- 给web服务和mysql-binlog服务添加DaemonSet 收集日志数据 

以上的三个服务都用Deployment来发布 ，因为Deployment是创建Pod和ReplicaSet，支持滚动升级的。  

### mysql用户名密码等信息可以放到Secret服务中。

## k8s 学习
StatefulSet 区分 于 Deployments和ReplicaSets  
StatefulSet 是为了解决有状态服务  
Deployments和ReplicaSets 是为无状态服务而设计

### service
可以理解为nginx upstream, 但service监听机器端口，获取数据后转发给pod处理。

### Ingress 
 可以理解为nginx server_name + location匹配， 匹配规则然后转发到相应的service
 
### PodPreset
pod公共环境变量，公共内容服务。 pod启动前从PodPreset获取信息。