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
		base.POST("/base/login", Login)
		// 单方法使用jwt中间件
		base.POST("/base/logout", extend.JWTAuth(), Logout)
		base.POST("/base/menu", extend.JWTAuth(), Menu)
		base.POST("/base/upload", extend.JWTAuth(), Upload)
	}
	
	api := route.Group("/api")
	api.Use()
	{
		api.POST("/gettopic", ApiGetTopic)
		api.POST("/getarticle", ApiGetArticle)
	}

	admin := route.Group("/admin")
	// 一组使用jwt验证
	admin.Use(extend.JWTAuth())
	{
		// 栏目
		admin.POST("/addtopic", AddTopic)
		admin.POST("/edittopic", EditTopic)
		admin.POST("/gettopic", GetTopic)
		admin.POST("/deltopic", DelTopic)

		// 文章
		admin.POST("/addarticle", AddArticle)
		admin.POST("/editarticle", EditArticle)
		admin.POST("/getarticle", GetArticle)
		admin.POST("/delarticle", DelArticle)
	}

	return route
}
