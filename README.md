## cde

基于Gin的快速API开发集成

## 简介

这是一个基于Gin框架搭建的快速API开发框架，旨在提供一套即开即用的API集成方案。该框架已经集成了常用的功能模块，如路由、参数绑定、错误处理、日志、db等，以便于使用者快速上手开发。

## 功能特点

- **基于Gin框架**：Gin是一款轻量级的Web框架，适用于快速构建高性能的API。
- **即开即用**：该框架已经集成了常用的功能模块，开箱即用，无需额外配置，直接敏捷开发。
- **统一错误处理**：通过中间件的方式，实现全局统一的错误处理和返回格式，避免重复代码。
- **日志记录**：使用zap库实现了高效、可定制的日志记录功能。
- **Swagger API文档**：通过集成Swagger，自动生成API文档，方便查看和测试API。
- **Database**: 通过gorm配置mysql作为数据库，redis做缓存。
- **Docker**: 使用Docker进行二进制部署

## 目录结构
```
├── handler                // 控制器层
│   ├── healthcheck.go     // 健康检查控制器
│   └── accountHandler.go  // 用户控制器示例
├── middleware        // 中间件层
│   ├── cors.go       // 跨域中间件
│   └── error.go      // 错误处理中间件
├── model             // 数据模型层
├── router            // 路由层
│   └── router.go     // 路由注册主文件
├── repository        // db 接口
├── util              // 工具函数目录
├── main.go           // 项目入口文件
└── README.md         // 项目说明文件
```
## 使用方法

1. 下载代码到本地：`git clone https://github.com/cde.git`
2. 安装依赖：`go mod tidy`
3. 参考`application_example.yaml`配置 `application.yaml`
4. 运行程序：`go run main.go`

Swagger API文档可以通过`http://localhost:8080/swagger/index.html`来访问。

## 持续开发
- [ ] 集成多个数据库
- [ ] 抽离部分服务，做微服务架构

## 贡献

如果你有好的想法或者发现了bug，欢迎提交issue或fork后提交PR。

## 授权许可

本项目采用MIT许可证。

