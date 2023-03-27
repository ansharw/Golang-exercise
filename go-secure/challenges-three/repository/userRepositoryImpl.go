package repository

import (
	"challenges-three/models"
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (repo *userRepository) FindByEmail(ctx context.Context, tx *gorm.DB, email string) string {
	User := models.User{}
	if err := tx.WithContext(ctx).Select("password").Where("email = ?", email).Take(&User).Error; err != nil {
		log.Fatalln("Error finding user password")
	}
	return User.Password
}

func (repo *userRepository) Login(ctx context.Context, tx *gorm.DB, email, pass string) (models.User, error) {
	User := models.User{}

	if err := tx.WithContext(ctx).Where("email = ? AND password = ?", email, pass).Take(&User).Error; err != nil {
		log.Fatalln("Error finding user")
	}

	return User, nil
}

func (repo *userRepository) Register(ctx context.Context, tx *gorm.DB, user models.User) (models.User, error) {
	User := models.User{}
	if err := tx.WithContext(ctx).Where("email = ?", user.Email).First(&User).Error; err == nil {
		log.Fatalln("Email Already Exists")
	}

	newUser := models.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := tx.WithContext(ctx).Create(&newUser).Error; err != nil {
		return user, fmt.Errorf("failed to register user with email %s: %w", user.Email, err)
	}

	return newUser, nil
}