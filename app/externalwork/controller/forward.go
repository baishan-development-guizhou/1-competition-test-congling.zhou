package controller

import (
	"1-competition-test-congling.zhou/app/externalwork/matedata/response"
	"github.com/gin-gonic/gin"
)

type forwardModel interface {
	// 转发
	ForwardRequests(c *gin.Context)
}

func (c *controller) ForwardRequests(ctx *gin.Context) {
	resp, err := c.service.ForwardRequests()
	if err != nil {
		response.Error(ctx, err, "内部错误了")
		return
	}
	ctx.JSON(200, &resp)
	return
}
