package services

import (
	"MyGram/helpers"
	"MyGram/model"
	"MyGram/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type userService struct {
	db       *gorm.DB
	repoUser repository.UserRepository
}

func NewUserService(db *gorm.DB, repoUser repository.UserRepository) *userService {
	return &userService{
		db:       db,
		repoUser: repoUser,
	}
}

func (service *userService) Login(ctx context.Context, user model.RequestUserLogin) (string, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if User, err := service.repoUser.Login(ctx, tx, user.Email, user.Password); err != nil {
		return "", err
	} else {
		token := helpers.GenerateToken(User.ID, user.Email)
		return token, nil
	}
}

func (service *userService) Register(ctx context.Context, user model.RequestUserRegister) (model.User, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	pass := helpers.HashPass(user.Password)
	user.Password = pass
	res, err := service.repoUser.Register(ctx, tx, user)
	if err != nil {
		return res, errors.New("register user canceled")
	} else {
		return res, nil
	}
}
