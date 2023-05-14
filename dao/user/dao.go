package user

import (
	"context"
	"game.interact.com/gin/frame/model"
)

type Dao interface {
	Info(ctx context.Context) *model.UserInfo
}

func NewDao() Dao {
	return newDao()
}
