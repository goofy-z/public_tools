### 工具简介
基于zap和lumberjack做了封装，支持日志按文件大小切割，支持console和json两种日志格式，支持企业微信告警。

### 使用方式
- 下载到本地
```go
go get https://git.code.oa.com/infosec_basicPlatform_devops/devops_sdk/zlog
```
- 项目内使用, 具体配置可查看本项目源码 ZapConfig定义的属性
```go
    git.code.oa.com/infosec_basicPlatform_devops/devops_sdk/zlog
    zcfg := zlog.NewZapConfig() 
    zcfg.LogDir = "xxxx"  // 默认是启动文件当前目录
    zcfg.CutSize = 3  // 默认是300M大小就切割文件
    zcfg.LoggerType = "console"   // json or console
    zcfg.ServerName = "NEW-TEST"  // 服务名称
    zcfg.Alarm = true  // 是否告警 默认否
    zcfg.AlarmUser = "ginxzheng"  // 告警user
    zcfg.AlarmGroup = "xxxx"  // 告警user
    zcfg.Build()  // 启动logger
    
    // 提供下面5种日志接口
    zlog.Errorf("测试Errorf")
    zlog.Error("测试Error")
    zlog.Infof("测试Infof")
    zlog.Info("测试Info")
    zlog.Fatal("测试Fatal")
```

### 效果
```shell script
2020-01-16 11:38:25.956	ERROR	[server: NEW-TEST]	zlog/zaplogger_test.go:15	测试Errorf
2020-01-16 11:38:27.978	ERROR	[server: NEW-TEST]	zlog/zaplogger_test.go:16	测试Error
2020-01-16 11:38:27.978	INFO	[server: NEW-TEST]	zlog/zaplogger_test.go:17	测试Infof
2020-01-16 11:38:27.978	INFO	[server: NEW-TEST]	zlog/zaplogger_test.go:18	测试Info
2020-01-16 11:38:28.037	FATAL	[server: NEW-TEST]	zlog/zaplogger_test.go:19	测试Fatal	{"stack": "git.code.oa.com/infosec_basicPlatform_devops/devops_sdk/zlog.(*consoleLogger).Fatal\n\t/Users/zheng/Desktop/gowork/sdkwork/src/devops_sdk/zlog/zapclog.go:73\ngit.code.oa.com/infosec_basicPlatform_devops/devops_sdk/zlog.Fatal\n\t/Users/zheng/Desktop/gowork/sdkwork/src/devops_sdk/zlog/zaplogger.go:92\ngit.code.oa.com/infosec_basicPlatform_devops/devops_sdk/zlog.TestLoggerFunction\n\t/Users/zheng/Desktop/gowork/sdkwork/src/devops_sdk/zlog/zaplogger_test.go:19\ntesting.tRunner\n\t/usr/local/go/src/testing/testing.go:909"}
```

###

     
