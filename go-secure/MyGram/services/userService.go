package services

import (
	"MyGram/model"
	"context"
)

type UserService interface {
	Login(ctx context.Context, user model.RequestUserLogin) (string, error)
	Register(ctx context.Context, user model.RequestUserRegister) (model.User, error)
}
