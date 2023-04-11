package repository

import (
	"MyGram/helpers"
	"MyGram/model"
	"context"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (repo *userRepository) Login(ctx context.Context, tx *gorm.DB, email, pass string) (model.User, error) {
	User := model.User{}

	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&User).Error; err != nil {
		log.Println("Error finding user")
		return User, errors.New("error finding user")
	}

	if err := helpers.ComparePass([]byte(User.Password), []byte(pass)); !err {
		return User, errors.New("error: User doesn't match")
	} else {
		return User, nil
	}
}

func (repo *userRepository) Register(ctx context.Context, tx *gorm.DB, user model.RequestUserRegister) (model.User, error) {
	User := model.User{}
	if err := tx.WithContext(ctx).Where("username = ? and email = ?", user.Username, user.Email).Take(&User).Error; err == nil {
		log.Println("Email and Username Already Exists")
		User.Email = ""
		User.Username = ""
		return User, errors.New("error: Email and Username Already Exists")
	} else if err := tx.WithContext(ctx).Where("username = ?", user.Username).Take(&User).Error; err == nil {
		User.Username = ""
		log.Println("Username Already Exists")
		return User, errors.New("error: Username Already Exists")
	} else if err := tx.WithContext(ctx).Where("email = ?", user.Email).Take(&User).Error; err == nil {
		User.Email = ""
		log.Println("Email Already Exists")
		return User, errors.New("error: Email Already Exists")
	}

	newUser := model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := tx.WithContext(ctx).Create(&newUser).Error; err != nil {
		return User, fmt.Errorf("failed to register user with email %s: %w", user.Email, err)
	}

	return newUser, nil
}
