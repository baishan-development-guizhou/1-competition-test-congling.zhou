package middleware

import (
	"time"

	"1-competition-test-congling.zhou/app/externalwork/domain/webEngin"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var _ Interface = (*middleware)(nil)

// Interface 中间件抽象接口
type Interface interface {
	Cors(ctx *gin.Context)
}

type middleware struct {
	webEngin webEngin.WebEngin
	cors     gin.HandlerFunc
}

// New 初始化路由
func New(webEngin webEngin.WebEngin) Interface {
	return &middleware{webEngin: webEngin,
		cors: cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "token", "authorization"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		})}
}
