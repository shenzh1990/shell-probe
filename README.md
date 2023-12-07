# 基于Go的服务器资源探针程序
## 介绍
这是使用Go语言开发的基于`gopsutil`获取服务器硬件资源信息的探针程序，支持两种通信方式：
- HTTP Get
- WebSocket

同时支持在Windows和Linux平台上运行，只需要运行响应平台的可执行文件即可。


> **注意**：程序默认启动将会监听端口，webapi：8082；
 
### 自定义启动
程序支持的命令如下：
```bash
-h 查看帮助信息，显示支持的命令
-v 查看程序版本信息
-V 查看程序版本信息
-fc 查看主机信息是否包含cpu 因为cpu接口比较缓慢 所以默认否 defaut[false]]
-ap 指定webapi服务监听的端口
-ta 指定网关访问地址 返回延迟信息 defaut[127.0.0.1:8082]
-token  启动设置接口调用的密码，webapi服务在请求头添加Authorization；websocket在请求URL添加token信息（127.0.0.1:8082/ws/host?token=123456）  defaut[123456] 
```
## 各服务接口说明
### WebApi服务接口
webapi服务启动之后，其地址为：`http://IP:port`
webapi服务提供如下路由：
- [GET]`/api/host`：获取主机所有硬件资源信息，包括：CPU、内存、硬盘、网络IO
- [GET]`/api/cpu`：获取主机CPU使用情况
- [GET]`/api/mem`：获取主机内存使用情况
- [GET]`/api/disk`：获取主机硬盘各分区使用情况
- [GET]`/api/net`：获取主机网络情况
- [GET]`/api/load`：获取主机网络情况
### websocket服务接口
websocket 服务提供如下路由：
- [GET]`/ws/host`：获取主机网络情况 升级为websocket形式，每3秒推送一次。
## 特别说明
- 由于gopsutil框架的实现方式，当获取CPU的使用情况时，需要等两秒钟才能够统计出CPU的使用情况，因此当通过webapi服务获取所有主机硬件资源使用情况时，需要等2秒才会有响应信息，若单独请求某一项硬件资源使用情况时，通常响应时间会在几毫秒。
- 代码源于https://gitee.com/mo_zhenshuang/server-resource-probe 设计开发
