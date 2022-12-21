@[TOC]
# go开发的论坛应用
## 技术知识点

1. gin框架
2. sqlx操作mysql
3. go-redis操作redis
4. viper读取配置文件并支持配置文件热更新
5. 雪花算法生成分布式uuid
6. zap处理应用日志，重写gin的looger和recovery处理框架的日志
7. jwt验证，和redis搭配实现限制同一时间一个账号仅一个设备登录
8. 使用validator做验证器和翻译器
9. md5加密
10. gin-swagger提供swagger接口文档
11. metrics提供符合prometheus处理的序列化数据
12. makefile编写和提供.air.conf，使用air进行热更新模式开发
13. 编写dockerfile和用于k8s的deployment方式部署的yaml

## 业务功能

1. 用户登录 注册 退出 用户信息
2. 帖子的社区频道增删改查
3. 帖子的增删改查
4. 帖子的投票

