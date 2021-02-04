package middleware

import (
	"github.com/gin-gonic/gin"
)

func (m *middleware) Cors(ctx *gin.Context) {
	m.cors(ctx)
	ctx.Next()
}
