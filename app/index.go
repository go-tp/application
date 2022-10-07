package app

import (
	"github.com/gin-gonic/gin"
	"gtp/model"
	"net/http"
)

// 首页
func Index(c *gin.Context){

	//c.String(200, "login")
	data := model.IndexM(c)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "请求成功",
			"data": data,
		})
	}
}

// 登录
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

// 登出
func Logout(c *gin.Context){
	c.String(200, "logout")
}

// 菜单
func Menu(c *gin.Context){
	c.String(200, "menu")
}

// 上传图片
func Upload(c *gin.Context){
	c.String(200, "upload")
}
