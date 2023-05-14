package user

import (
	"context"
	"game.interact.com/gin/frame/gerror"
	"game.interact.com/gin/frame/model"
	"game.interact.com/gin/frame/service/user"
	"game.interact.com/gin/frame/util/decorator"
)

type handler struct {
	userSrv user.Service
}

func newHandler() *handler {
	return &handler{
		userSrv: user.NewService(),
	}
}

func (h *handler) Info(c *decorator.RequestContext) (model.Resp, gerror.Error) {
	params := c.Params().(*model.UserInfoReq)
	if params == nil {
		//logs.CtxError(ctx, "ctx.Params().(*businessproto.PayOrderRequest) err")
		return nil, gerror.RequestParamsError
	}

	info := h.userSrv.Info(context.Background())
	if params.ID == 1 {
		return nil, gerror.ServiceUnavailable.WithCustomMessage("err:id=1")
	}

	return &model.UserInfoResp{
		Name: info.Name,
		Age:  info.Age + params.ID,
	}, nil

}
