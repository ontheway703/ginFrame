package user

import (
	"context"
	"game.interact.com/gin/frame/dao/user"
	"game.interact.com/gin/frame/model"
)

type srv struct {
	userDao user.Dao
}

func newService() *srv {
	return &srv{
		userDao: user.NewDao(),
	}
}

func (s *srv) Info(ctx context.Context) *model.UserInfo {
	return s.userDao.Info(ctx)
}
