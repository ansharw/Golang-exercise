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
		log.Fatalln("Error finding user")
		return User, errors.New("Error finding user")
	}

	if err := helpers.ComparePass([]byte(User.Password), []byte(pass)); !err {
		return User, errors.New("error: User doesn't match")
	} else {
		return User, nil
	}
}

func (repo *userRepository) Register(ctx context.Context, tx *gorm.DB, user model.User) (model.User, error) {
	User := model.User{}
	if err := tx.WithContext(ctx).Where("email = ?", user.Email).Take(&User).Error; err == nil {
		log.Fatalln("Email Already Exists")
		return User, errors.New("error: Email Already Exists")
	}

	newUser := model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := tx.WithContext(ctx).Create(&newUser).Error; err != nil {
		return user, fmt.Errorf("failed to register user with email %s: %w", user.Email, err)
	}

	return newUser, nil
}
