package response

import (
	"1-competition-test-congling.zhou/app/externalwork/matedata/code"
	"1-competition-test-congling.zhou/app/externalwork/matedata/error_msg"
	"github.com/gin-gonic/gin"
)

func Res(ctx *gin.Context, res *code.Result) {
	ctx.JSON(200, &res)
}

func OK(ctx *gin.Context, data interface{}) {
	Res(ctx, &code.Result{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}

func Error(ctx *gin.Context, err error, msg ...string) {
	if err == nil {
		OK(ctx, nil)
		return
	}
	var message = error_msg.InternalError
	if len(msg) > 0 {
		message = msg[0]
	}

	Res(ctx, &code.Result{
		Code: 1,
		Msg:  message,
		Data: nil,
	})
	return
}
