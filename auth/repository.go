package auth

import (
	"context"
	"github.com/IDontCareMe/bookmark_REST_API/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
}
