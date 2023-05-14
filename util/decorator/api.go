package decorator

import (
	"encoding/json"
	"game.interact.com/gin/frame/gerror"
	"game.interact.com/gin/frame/model"
	"github.com/gin-gonic/gin"
	"reflect"
)

type RequestHandler func(context *RequestContext) (model.Resp, gerror.Error)

func ApiDecorator(handler RequestHandler, paramsMsg model.Req) func(*gin.Context) {
	return func(c *gin.Context) {
		//构造新请求
		ctx := NewRequestContext(c)

		// 解析参数
		if paramsMsg != nil {
			err := ctx.InitRequestParams(reflect.TypeOf(paramsMsg).Elem())
			if err != nil {
				genJsonErrResp(ctx, err)
				return
			}
		}

		// 执行 handler
		respMsg, bizErr := handler(ctx)
		if bizErr != nil {
			genJsonErrResp(ctx, bizErr)
			return
		}
		if respMsg == nil {
			respMsg = &model.EmptyEsp{}
		}

		b, err := json.Marshal(respMsg)
		if err != nil {
			genJsonErrResp(ctx, gerror.ServiceUnavailable)
			return
		}
		genJsonResp(ctx, string(b))
		return
	}
}

func genJsonErrResp(ctx *RequestContext, err gerror.Error) {
	ctx.GinContext.JSON(200, gin.H{
		"message":     err.Error(),
		"status_code": err.StatusCode(),
		"data":        "",
	})
}

func genJsonResp(ctx *RequestContext, data string) {
	ctx.GinContext.JSON(200, gin.H{
		"message":     "",
		"status_code": 0,
		"data":        data,
	})
}
