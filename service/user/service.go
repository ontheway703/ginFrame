package user

import (
	"context"
	"game.interact.com/gin/frame/model"
)

type Service interface {
	Info(ctx context.Context) *model.UserInfo
}

func NewService() Service {
	return newService()
}
