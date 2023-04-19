# go-template
---
```go-template```是基于Golang的API模板项目，提供了一些常用的功能Demo，以供参考学习！

## 目录结构
```yaml
go-template:
    - assets
    - bin
    - cmd
    - conf
    - deployment
    - docs
    - global
    - internal
    - log
    - middleware
    - pkg
    - router
    - script
```

```text
assets:     存放资源文件
bin:        存放二进制文件
conf:       存放配置文件
cmd:        存放程序入口代码
deployment: 存放部署相关代码
docs:       存放文档
global:     存放全局变量
internal:   存放业务代码
log:        存放日志文件
middleware: 存放中间件代码
pkg:        存放模型代码
router:     存放路由代码
runtime:    存放运行时生成文件
script:     存放脚本文件
```

## 集成组件
1. 支持 [viper](https://github.com/spf13/viper) 组件，用以解析配置文件
2. 支持 [gorm](https://gorm.io/gorm) 组件，用以连接数据库
3. 支持 [zap](https://go.uber.org/zap) 组件，用以收集日志
4. 支持 RESTFUL，用以规范接口
5. 支持 [Swagger](https://github.com/swaggo/gin-swagger) 组件，用以生成接口文档
6. 支持 [websocket](https://github.com/gorilla/websocket)，实现实时通讯
7. 支持 [cron](https://github.com/jakecoffman/cron)， 实现定时任务
8. 支持 [jwt](https://github.com/golang-jwt/jwt) 组件，实现权限管理

## Acknowledgments
以下项目对```go-template```有重大参考意义
- [mix-go/mix](https://github.com/mix-go/mix) ✨ Standard Toolkit for Go fast development / Go 快速开发标准工具包
- [JuneZhuSummer/go-api](https://github.com/JuneZhuSummer/go-api) 这是一个go的api模板项目
- [flipped-aurora/gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin) 基于vite+vue3+gin搭建的开发基础平台，集成多项开发必备功能。

## 联系作者
![联系作者](https://github.com/feiria/go-template/assets/img.png)