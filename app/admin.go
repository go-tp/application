package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gtp/model"
)

// 新增栏目
func AddTopic(c *gin.Context){
	data := model.AddTopicM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}

// 编辑栏目
func EditTopic(c *gin.Context){
	data := model.EditTopicM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}

// 读取栏目
func GetTopic(c *gin.Context){
	data := model.GetTopicM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}

// 删除栏目
func DelTopic(c *gin.Context){
	data := model.DelTopicM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}

// 新增文章
func AddArticle(c *gin.Context){
	data := model.AddArticleM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}

// 编辑文章
func EditArticle(c *gin.Context){
	data := model.AddArticleM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}

// 读取文章
func GetArticle(c *gin.Context){
	data := model.AddArticleM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}

// 读取文章列表
func GetArticleList(c *gin.Context){
	data := model.AddArticleM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}


// 删除文章
func DelArticle(c *gin.Context){
	data := model.AddArticleM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}