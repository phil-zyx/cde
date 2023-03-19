package router

import (
	docs "github.com/cde/docs"
	"github.com/cde/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"
	// 创建一个默认路由
	r := gin.Default()
	routers(r)
	return r
}

func routers(r *gin.Engine) {
	api := r.Group("/api")
	{
		// v1
		v1 := api.Group("/v1")
		// 登录|注册
		authGroup := v1.Group("/account")
		{
			authGroup.POST("/register", handler.Register)
			authGroup.POST("/login", handler.Login)
		}
		//// 分类
		//categoryGroup := v1.Group("/category", util.CheckToken())
		//{
		//	categoryGroup.GET("/lists", handler.GetCategoryLists)
		//	categoryGroup.GET("/goods-lists", handler.GetCategoryGoodsLists)
		//}
		//// 商品
		//goodsGroup := v1.Group("/goods", middleware.CheckToken())
		//{
		//	goodsGroup.GET("/detail", handler.GoodsDetails)
		//}
		//// 用户
		//userGroup := v1.Group("/user", middleware.CheckToken())
		//{
		//	userGroup.POST("/bind-phone", handler.UserBindPhone)
		//	userGroup.POST("/un-bind-phone", handler.UnUserBindPhone)
		//}
	}
}
