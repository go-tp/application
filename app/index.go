package app

import (
	"github.com/gin-gonic/gin"
	"gtp/model"
	"net/http"
)

func Index(c *gin.Context){
	data := model.IndexM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}

func Login(c *gin.Context){
	c.String(200, "login")
}

func Logout(c *gin.Context){
	c.String(200, "logout")
}

func Menu(c *gin.Context){
	c.String(200, "menu")
}
