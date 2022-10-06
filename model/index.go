package model

import (
	"github.com/gin-gonic/gin"
	"gtp/extend"
	"fmt"
	"github.com/garyburd/redigo/redis"
)


type index struct {
	Id int `form:"id" json:"id"`
}

func IndexM(c *gin.Context) interface{} {
	// 连接 mysql
	mysqlCon,_ := extend.MysqlInit()
	// 连接 redis
	extend.RedisInit()

	//c.String(200, Db)

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
	fmt.Printf("type=> %T\n", v)
	// unit8
	fmt.Println(v)
	
	// string
	fmt.Println(v1)
	


	
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


	//r := index{}
	return v1
}


func LoginM(c *gin.Context) interface{} {
	// jwt
	// 中间件需要验证jwt ParseToken
	// ctx := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDczMjQ5MjMsInVpZCI6IjEifQ.ooWMRCN_ZVvHxKMX3OnuJxviYEVUI8xCScaM3M0bRXY"

	// claims,_ := extend.ParseToken(ctx);
	// fmt.Println(claims)

	// 此方法需要根据用户登录情况 创建token
	UserId := "1"
	// 生成jwt-token
	token,_ := extend.GenerateToken(UserId)
	return token
}
