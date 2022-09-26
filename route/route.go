package route

import (
	"github.com/gin-gonic/gin"
	."gtp/app"
)

func Init() *gin.Engine {
	var route *gin.Engine

	route = gin.Default()

	base := route.Group("/")
	base.Use()
	{
		base.GET("/", Index)
		base.POST("/login", Login)
		base.POST("/logout", Logout)
		base.POST("/menu", Menu)
	}
	
	apiArticle := route.Group("/api")
	apiArticle.Use()
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