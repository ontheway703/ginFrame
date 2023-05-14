package user

import (
	"game.interact.com/gin/frame/gerror"
	"game.interact.com/gin/frame/model"
	"game.interact.com/gin/frame/util/decorator"
)

type Handler interface {
	Info(c *decorator.RequestContext) (model.Resp, gerror.Error)
}

func NewHandler() Handler {
	return newHandler()
}
