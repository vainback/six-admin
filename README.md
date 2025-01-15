# Six-Admin多租户中后台管理系统

```
采用MIT协议，商业化完全免费且无需保留任何商标，如需采用，感谢您留下宝贵的Star
```


## 一、预览(演示)地址

[演示地址](https://dl.sorks.cn/admin)

文档暂缺，正在编辑中

## 二、技术栈

### 版本要求

``` go

Go Version >= 1.23
至少保证 Go 1.23 以上 需开启环境变量 GOEXPERIMENT=aliastypeparams
或更新至 Go 1.24 以上
设置 go mod 代理为：GOPROXY=https://goproxy.cn,direct

node version >= 18.0
npm建议安装nrm包 随时切换代理


已对 COS OSS 七牛云 进行简单适配
```

### 后端技术栈
[1. Gin - Web 框架](https://gin-gonic.com/zh-cn/docs/)

[2. Gorm - 数据库操作ORM框架](https://gorm.io/zh_CN/docs/)

[3. ozzo-validation - 验证器](https://github.com/go-ozzo/ozzo-validation)

[4. go-redis - Redis链接库](https://github.com/redis/go-redis/v9)

[5. robfig-cron - 定时任务](https://github.com/robfig/cron/v3)

[6. viper - 配置文件解析库](https://github.com/spf13/viper)

[7. cast - 类型转换库](https://github.com/spf13/cast)

### 前端UI框架
[Arco Design Pro Vue](https://arco.design/vue/docs/pro/start)


```
缺点：未对国际化做兼容
```


