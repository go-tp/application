package app

import (
	"github.com/gin-gonic/gin"
)

// 新增栏目
func AddTopic(c *gin.Context){
	c.String(200, "AddTopic")
}

// 编辑栏目
func EditTopic(c *gin.Context){
	c.String(200, "EditTopic")
}

// 读取栏目
func GetTopic(c *gin.Context){
	c.String(200, "GetTopic")
}

// 删除栏目
func DelTopic(c *gin.Context){
	c.String(200, "DelTopic")
}

// 新增文章
func AddArticle(c *gin.Context){
	c.String(200, "AddArticle")
}

// 编辑文章
func EditArticle(c *gin.Context){
	c.String(200, "EditArticle")
}

// 读取文章
func GetArticle(c *gin.Context){
	c.String(200, "GetArticle")
}

// 删除文章
func DelArticle(c *gin.Context){
	c.String(200, "DelArticle")
}