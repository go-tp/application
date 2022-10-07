package app

import (
	"github.com/gin-gonic/gin"
)

// 读取栏目
func ApiGetTopic(c *gin.Context){
	c.String(200, "ApiGetTopic")
}

// 读取文章
func ApiGetArticle(c *gin.Context){
	c.String(200, "ApiGetArticle")
}


