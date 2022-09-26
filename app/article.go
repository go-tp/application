package app

import (
	"github.com/gin-gonic/gin"
)

func ApiGetArtList(c *gin.Context){
	c.String(200, "ApiGetArtList")
}


func ApiGetTopic(c *gin.Context){
	c.String(200, "ApiGetTopic")
}