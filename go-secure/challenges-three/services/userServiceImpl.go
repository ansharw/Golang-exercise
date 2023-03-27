package services

import (
	"challenges-three/helpers"
	"challenges-three/models"
	"challenges-three/repository"
	"context"
	"errors"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type userService struct {
	db        *gorm.DB
	repoUser  repository.UserRepository
	validator *validator.Validate
}

func NewUserService(db *gorm.DB, repoUser repository.UserRepository, validator_ validator.Validate) *userService {
	return &userService{
		db:        db,
		repoUser:  repoUser,
		validator: &validator_,
	}
}

func (service *userService) Login(ctx context.Context, user models.User) (string, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	passHashed := service.repoUser.FindByEmail(ctx, tx, user.Email)
	err := helpers.ComparePass([]byte(passHashed), []byte(user.Password))
	if !err {
		return "", errors.New("invalid email/password")
	} else {
		User, _ := service.repoUser.Login(ctx, tx, user.Email, passHashed)
		token := helpers.GenerateToken(User.ID, user.Email)
		return token, nil
	}
}

func (service *userService) Register(ctx context.Context, user models.User) (models.User, error) {
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