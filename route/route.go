package route

import (
	"github.com/gin-gonic/gin"
	."gtp/app"
	"gtp/extend"
)

func Init() *gin.Engine {
	var route *gin.Engine

	route = gin.Default()

	base := route.Group("/")
	base.Use()
	{
		base.GET("/", Index)
		base.POST("/login", Login)
		// 单方法使用jwt中间件
		base.POST("/logout", extend.JWTAuth(),Logout)
		base.POST("/menu", Menu)
	}
	
	apiArticle := route.Group("/api")
	// 一组使用jwt验证
	apiArticle.Use(extend.JWTAuth())
	{
		apiArticle.POST("/article/getTopic", ApiGetArtList)
		apiArticle.POST("/article/getArtList", ApiGetTopic)
	}

	adminArticle := route.Group("/admin")
	adminArticle.Use()
	{
		adminArticle.POST("/article/getTopic", AdminGetArtTopic)
		adminArticle.POST("/article/getArtList", AdminGetArtList)
	}

	return route
}
