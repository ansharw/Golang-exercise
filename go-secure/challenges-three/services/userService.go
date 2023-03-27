package services

import (
	"challenges-three/models"
	"context"
)

type UserService interface {
	Login(ctx context.Context, user models.User) (string, error)
	Register(ctx context.Context, user models.User) (models.User, error)
}
