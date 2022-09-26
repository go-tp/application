package app

import (
	"github.com/gin-gonic/gin"
)

func AdminGetArtList(c *gin.Context){
	c.String(200, "AdminGetArtList")
}


func AdminGetArtTopic(c *gin.Context){
	c.String(200, "AdminGetArtTopic")
}