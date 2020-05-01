### 工具简介
封装于基础库 http，做了个性化改动，依赖zlog日志

### 使用方式
- 下载到本地
```go
go get https://git.code.oa.com/infosec_basicPlatform_devops/devops_sdk/httpclient
```
- 项目内使用，返回的res也是 *http.Response, 基本和原生http库一致，多了重试，日志收集, 配置代理
```go
    git.code.oa.com/infosec_basicPlatform_devops/devops_sdk/httpclient
    // 根据options文件提供的方法设置基础属性
    client := httpclient.NewClient(httpclient.SetTimeout(3))
    MonitorURL := "http://monitor-api.server.com/get_day_data"
    res ,err := client.GET(MonitorURL, headers)  // POST PUT DELETE
    defer res.Body.Close()
```

     
