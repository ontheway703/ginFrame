package user

import (
	"context"
	"game.interact.com/gin/frame/model"
)

type dao struct {
}

func newDao() *dao {
	return &dao{}
}

func (d *dao) Info(ctx context.Context) *model.UserInfo {

	return &model.UserInfo{
		Name: "ll",
		Age:  22,
	}

}
