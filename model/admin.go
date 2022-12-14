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
	fmt.Println("username:",username)

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
	defer mysqlCon.Close()
	return id
}

type queryEditTopic struct {
	Title string `form:"title" json:"title"`
	Sort int `form:"sort" json:"sort"`
	Id int `form:"id" json:"id"`
}

func EditTopicM(c *gin.Context) interface{} {
	username,_ := c.Get("username")
	fmt.Println("username:",username)

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
	defer mysqlCon.Close()
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
	fmt.Println("username:",username)

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

type queryAddArticle struct {
	Title string `form:"title" json:"title"`
	Litpic string `form:"litpic" json:"litpic"`
	Content string `form:"content" json:"content"`
	Topic_id int `form:"topic_id" json:"topic_id"`
}

func AddArticleM(c *gin.Context) interface{} {
	username,_ := c.Get("username")
	fmt.Println("username:",username)

	// 获取参数
	query := queryAddArticle{}
	c.ShouldBind(&query)
	title := query.Title
	litpic := query.Litpic
	content := query.Content
	topic_id := query.Topic_id
	fmt.Println(content)
	if topic_id<=0{ return "topic_id为空"}
	// 连接 mysql
	mysqlCon,_ := extend.MysqlInit()
	// 连接 redis
	extend.RedisInit()
	now := extend.Now()
	// 执行插入
	result, _ := mysqlCon.Exec("INSERT INTO article (title,status,created_at,updated_at,litpic,content,topic_id,user_id) VALUES (?,?,?,?,?,?,?,?)",title,"1",now,now,litpic,content,topic_id,username)
	// // 获取插入id
	id, _ := result.LastInsertId()
	defer mysqlCon.Close()
	return id
}

type queryEditArticle struct {
	Title string `form:"title" json:"title"`
	Litpic string `form:"litpic" json:"litpic"`
	Content string `form:"content" json:"content"`
	Topic_id int `form:"topic_id" json:"topic_id"`
	Id int `form:"id" json:"id"`
}

type dbarticle struct{
	title_ *string `db:"title" json:",omitempty"`
	litpic_ *string `db:"litpic" json:",omitempty"`
	content_ *string `db:"content" json:",omitempty"`
	topic_id_ *int `db:"topic_id" json:",omitempty"`
}

func EditArticleM(c *gin.Context) interface{} {
	username,_ := c.Get("username")
	fmt.Println("username:",username)

	// 获取参数
	query := queryEditArticle{}
	c.ShouldBind(&query)
	title := query.Title
	litpic := query.Litpic
	content := query.Content
	topic_id := query.Topic_id
	id := query.Id

	if id <= 0 { return "请输入id"}
	
	// 连接 mysql
	mysqlCon,_ := extend.MysqlInit()
	// 连接 redis
	extend.RedisInit()

	var dbarticle dbarticle
	mysqlCon.QueryRow("SELECT title,litpic,content,topic_id FROM article WHERE id=?", id).Scan(
		&dbarticle.title_,
		&dbarticle.litpic_,
		&dbarticle.content_,
		&dbarticle.topic_id_)

	if title=="" { title = *dbarticle.title_}
	if litpic=="" { litpic = *dbarticle.litpic_}
	if content=="" { content = *dbarticle.content_}
	if topic_id==0 { topic_id = *dbarticle.topic_id_}

	now := extend.Now()
	
	// 执行插入
	result, _ := mysqlCon.Exec("UPDATE article SET title=?,litpic=?,content=?,topic_id=?,updated_at=? where status = 1 and id = ?",title,litpic,content,topic_id,now,id)
	// 影响行数
	n,_ := result.RowsAffected()
	defer mysqlCon.Close()
	return n
}

type articlelist struct {
	id   int
	title  string
	topic_id int
	litpic *string `db:"litpic" json:",omitempty"`
	status int
	created_at int
	updated_at int
}

type queryGetArticleList struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

func GetArticleListM(c *gin.Context) interface{} {
	// 获取参数
	query := queryGetArticleList{}
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
	sqlStr := "SELECT id, title, topic_id,status,created_at,updated_at,litpic FROM article WHERE status > 0 order by created_at desc limit ?,?"
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
	var t articlelist
	for rows.Next() {
		t1 := make(map[string]interface{})
		rows.Scan(&t.id, &t.title, &t.topic_id, &t.status,&t.created_at,&t.updated_at,&t.litpic)
		t1["id"] = t.id
		t1["title"] = t.title
		t1["topic_id"] = t.topic_id
		t1["status"] = t.status
		t1["litpic"] = t.litpic
		t1["created_at"] = extend.Date(t.created_at)
		t1["updated_at"] = extend.Date(t.updated_at)
		result_data = append(result_data, t1)
	}

	return result_data
}

type article struct {
	id   int
	title  string
	topic_id int
	litpic *string `db:"litpic" json:",omitempty"`
	content *string `db:"content" json:",omitempty"`
	status int
	created_at int
	updated_at int
}

type queryGetArticle struct {
	Id int `form:"id" json:"id"`
}

func GetArticleM(c *gin.Context) interface{} {
	// 获取参数
	query := queryGetArticle{}
	c.ShouldBind(&query)
	id := query.Id

	if id <=0 { return "id不能为空"}

	db,_ := extend.MysqlInit()
	// todo 增加分页
	sqlStr := "SELECT id, title, litpic,status,created_at,updated_at,content FROM article WHERE status > 0 and id = ?"
	var t article
	db.QueryRow(sqlStr,id).Scan(
		&t.id, 
		&t.title, 
		&t.litpic, 
		&t.status,
		&t.created_at,
		&t.updated_at,
		&t.content)
	// 重要：关闭 rows, 释放持有的数据库链接
	defer db.Close()
	
	t1 := make(map[string]interface{})
	t1["id"] = t.id
	t1["title"] = t.title
	t1["litpic"] = t.litpic
	t1["status"] = t.status
	t1["content"] = t.content
	t1["created_at"] = extend.Date(t.created_at)
	t1["updated_at"] = extend.Date(t.updated_at)

	return t1
}

type queryDelArticle struct {
	Id int `form:"id" json:"id"`
}

func DelArticleM(c *gin.Context) interface{} {
	username,_ := c.Get("username")
	fmt.Println("username:",username)

	// 获取参数
	query := queryDelArticle{}
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
	result, _ := mysqlCon.Exec("UPDATE article SET status = 0,updated_at=? where status = 1 and id = ?",now,id)
	defer mysqlCon.Close()
	// 影响行数
	n,_ := result.RowsAffected()
	return n
}