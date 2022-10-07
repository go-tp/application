package model

import (
	"github.com/gin-gonic/gin"
	"gtp/extend"
	"fmt"
	_"github.com/garyburd/redigo/redis"
)

type queryAddTopic struct {
	Title string `form:"title" json:"title"`
	Sort int `form:"sort" json:"sort"`
}

func AddTopicM(c *gin.Context) interface{} {
	username,_ := c.Get("username")
	fmt.Println(username)

	// 获取参数
	query := queryAddTopic{}
	c.ShouldBind(&query)
	title := query.Title
	sort := query.Sort

	if sort<=0{ sort=1}

	// 连接 mysql
	mysqlCon,_ := extend.MysqlInit()
	// 连接 redis
	extend.RedisInit()
	now := extend.Now()
	// 执行插入
	result, _ := mysqlCon.Exec("INSERT INTO topic (title,status,created_at,updated_at,sort) VALUES (?,?,?,?,?)",title,"1",now,now,sort)
	// 获取插入id
	id, _ := result.LastInsertId()

	return id
}

type queryEditTopic struct {
	Title string `form:"title" json:"title"`
	Sort int `form:"sort" json:"sort"`
	Id int `form:"id" json:"id"`
}

func EditTopicM(c *gin.Context) interface{} {
	username,_ := c.Get("username")
	fmt.Println(username)

	// 获取参数
	query := queryEditTopic{}
	c.ShouldBind(&query)
	title := query.Title
	sort := query.Sort
	id := query.Id

	if id <= 0 { return "请输入id"}
	
	// 连接 mysql
	mysqlCon,_ := extend.MysqlInit()
	// 连接 redis
	extend.RedisInit()

	if sort<=0 {
		// 取原值
		mysqlCon.QueryRow("SELECT sort FROM topic WHERE id=?", id).Scan(&sort)
	}
	now := extend.Now()
	
	// 执行插入
	result, _ := mysqlCon.Exec("UPDATE topic SET title=?,sort=?,updated_at=? where status = 1 and id = ?",title,sort,now,id)
	// 影响行数
	n,_ := result.RowsAffected()
	return n
}

type topic struct {
	id   int
	title  string
	sort int
	status int
	created_at int
	updated_at int
}

type queryGetTopic struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

func GetTopicM(c *gin.Context) interface{} {
	// 获取参数
	query := queryGetTopic{}
	c.ShouldBind(&query)
	page := query.Page
	size := query.Size

	// 默认值
	if page <=0 { page = 1}
	if size <=0 { size = 10}

	// 分页
	start,end := extend.Page(page,size)

	db,_ := extend.MysqlInit()
	// todo 增加分页
	sqlStr := "SELECT id, title, sort,status,created_at,updated_at FROM topic WHERE status > 0 order by created_at desc,sort desc limit ?,?"
	rows,err := db.Query(sqlStr,start,end)
	if err != nil {
		fmt.Printf("查询失败, err:%v\n", err)
		// defer rows.Close()
		return ""
	}
	// 重要：关闭 rows, 释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	result_data := make([]map[string]interface{}, 0)
	var t topic
	for rows.Next() {
		t1 := make(map[string]interface{})
		rows.Scan(&t.id, &t.title, &t.sort, &t.status,&t.created_at,&t.updated_at)
		t1["id"] = t.id
		t1["title"] = t.title
		t1["sort"] = t.sort
		t1["status"] = t.status
		t1["created_at"] = extend.Date(t.created_at)
		t1["updated_at"] = extend.Date(t.updated_at)
		result_data = append(result_data, t1)
	}

	return result_data
}

type queryDelTopic struct {
	Id int `form:"id" json:"id"`
}

func DelTopicM(c *gin.Context) interface{} {
	username,_ := c.Get("username")
	fmt.Println(username)

	// 获取参数
	query := queryDelTopic{}
	c.ShouldBind(&query)
	id := query.Id

	if id <= 0{
		return "请输入id"
	}
	// 时间戳
	now := extend.Now()
	// 连接 mysql
	mysqlCon,_ := extend.MysqlInit()
	// 连接 redis
	extend.RedisInit()
	// 执行插入
	result, _ := mysqlCon.Exec("UPDATE topic SET status = 0,updated_at=? where status = 1 and id = ?",now,id)
	defer mysqlCon.Close()
	// 影响行数
	n,_ := result.RowsAffected()
	return n
}