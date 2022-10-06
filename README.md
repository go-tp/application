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

# 热更新[go install github.com/pilu/fresh@latest]
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

## 路由 [route/route.go]
```
base := route.Group("/")
base.Use()
{
    base.GET("/", Index)
    base.POST("/login", Login)
    base.POST("/logout", Logout)
    base.POST("/menu", Menu)
}
```

## 控制中心 [app/*]
```
func Login(c *gin.Context){
	
	data := model.LoginM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}
```
## 模型 [model/*]
```
func LoginM(c *gin.Context) interface{} {
	return "login"
}
```
# jwt

> 中间件使用
```
# route/route.go
import "gtp/extend"

route.Use(extend.JWTAuth())

或 在单一路由下使用

route.POST("/logout", extend.JWTAuth(),Logout)

```

> 加密方法
```
func GenerateToken(user string) (string, error)
```

> 解密方法
```
func ParseToken(tokenStr string) (jwt.MapClaims, error)
```

# 数据库连接
```
// 引入扩展
import "gtp/extend"

// 连接 mysql
mysqlCon,_ := extend.MysqlInit()

// mysql

result, _ := mysqlCon.Exec("INSERT INTO article (title, content, status,created_at,updated_at) VALUES (?,?,?, ?, ?)","文章1","正文1","1","1664378735","1664378735")

fmt.Println("result:",  result)


// 插入单条
stmt, _ := mysqlCon.Prepare("INSERT article SET title=?,content=?,status=?,created_at=?,updated_at=?")
res, _ := stmt.Exec("文章1","正文1","1","1664378735","1664378735")

id, _ := res.LastInsertId()

fmt.Println("id:",      id)

// 查询单条
var title, content string
var status int
mysqlCon.QueryRow("SELECT title,content,status FROM article WHERE id=?", 1).Scan(&title, &content, &status)

fmt.Println("title:",   title)
fmt.Println("content:", content)
fmt.Println("status:",  status)

// 查询多条

// 事务
tx, err := mysqlCon.Begin()
fmt.Println("err:",     err)

err1 := tx.Commit()
fmt.Println("err:",     err1)

err2 := tx.Rollback()
fmt.Println("err:",     err2)
```

# redis
```
// 引入扩展
import "gtp/extend"

// 初始化连接池
extend.RedisInit()

// redis 使用
r_key := "indexapi:getindex"
rdb := extend.Pool
conn := rdb.Get()
defer conn.Close()
// redis 设置
conn.Do("Set", r_key,"2")
// redis 获取
v,_ := conn.Do("GET",r_key)

v1,_ := redis.String(v,nil)
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