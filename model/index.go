package model

import (
	"github.com/gin-gonic/gin"
	"gtp/extend"
	_"fmt"
	_"github.com/garyburd/redigo/redis"
)

func IndexM(c *gin.Context) interface{} {
	return "Welcome to go-tp.com."
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

// 上传图片
func UploadM(c *gin.Context) interface{} {
    file, errLoad := c.FormFile("file")
    if errLoad != nil {
        msg := "获取上传文件错误"
        return msg
    }
 
    // 上传文件到指定的路径
    ret := make(map[string]string)
	// 名称确保唯一
	uuid := extend.Uuid()
    ret["fileName"] = uuid + file.Filename
    ret["fileNameOrigin"] = file.Filename
	
    // public/upload 文件夹下
    filePath := "./public/upload/"+ ret["fileName"]

    err := c.SaveUploadedFile(file, filePath)
    if err != nil {
        return err
    }
    
    ret["picUrl"] = ret["fileName"]
	return ret["picUrl"]
}

func LogoutM(c *gin.Context) interface{} {
	// 清除缓存
	// 清除redis
	return "退出成功."
}

func MenuM(c *gin.Context) interface{} {
	return "菜单."
}