# douyin

learned from: https://github.com/bytedance-youthcamp-jbzx/tiktok

项目目录

kitex: 和前端的数据定义

internal: 内部rpc的数据定义

cmd/api/handler: 接前端请求，调rpc

cmd/api/rpc: 接rpc请求，调业务逻辑

cmd/api/xxx: 业务逻辑，微服务

pkg: 对第三方库的封装

dal: 数据库操作

config: 配置文件

scripts: 启动脚本