# go-tp

> 拉取代码
```
git clone https://github.com/go-tp/application
或
git clone https://github.com/go-tp/application gtp [项目名]
```

> 修改配置文件 
>>具体参数根据本地配置进行修改
```
config.yaml.bak 修改为 config.yaml
```


>快速开始
```
# 更新依赖
go mod tidy
go mod vendor

# 热更新[go install github.com/pilu/fresh@laster]
fresh 
```

>编译或打包
```
go run main.go
go build main.go
```

>目录结构
```
项目名
├─app                controller应用目录
│  ├─admin.go        后台接口
│  ├─article.go      api接口详情
│  │─index.go        首页模块
│  └─ ...  
│
├─configs                应用配置目录
│  └─conf.go            读取配置文件
│
├─model                 模型目录
│  └─index.go           模型文件
│  
├─docker                 docker目录
│  └─...
│
├─route                 路由定义目录
│  ├─route.go           路由定义
│  └─...                更多
│
├─public                WEB目录（对外访问目录）
│  ├─css                css
│  ├─js                 js
│  └─...                更多                  
│
├─extend                扩展类库目录
├─vendor                第三方类库目录（go get依赖库）
├─config.yaml.bak       配置文件，记得根据自身情况修改，后期进行优化
├─go.mod                go get 定义文件
├─LICENSE               MIT授权说明文件
└─README.md             README 文件
```


# 问题

>解决inotify_init1() failed: Too many open files

```
[root@go-tp]# vim  /etc/sysctl.conf

添加
fs.file-max=9000000
fs.inotify.max_user_instances = 1000000
fs.inotify.max_user_watches = 1000000


sysctl -p
```


> # golang.org/x/net/http2
>vendor/golang.org/x/net/http2/transport.go:416:45: undefined: os.ErrDeadlineExceeded


```
使用 go 1.17以上版本
```