package route

import (
	"1-competition-test-congling.zhou/app/externalwork/controller"
	"1-competition-test-congling.zhou/app/externalwork/middleware"
	"github.com/gin-gonic/gin"
)

// Route 路由实例
type Route struct {
	controller controller.Interface
	middleware middleware.Interface
}

// New 初始化路由
func New(controller controller.Interface, middleware middleware.Interface) *Route {
	return &Route{
		controller: controller,
		middleware: middleware,
	}
}

// Install 路由初始化
func (r *Route) Install(ginEngin *gin.Engine) {
	// 全局中间件加载
	ginEngin.Use()

	now := ginEngin.Group("/now")
	{
		now.POST("/ggg")
	}
}
