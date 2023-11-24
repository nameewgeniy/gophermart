package repositories

import (
	"context"
	"github.com/google/uuid"
	"gophermart/internal/domain/models"
)

type Storage interface {
	Up(ctx context.Context) error
	Down(ctx context.Context) error
}

type UserRepository interface {
	CreateUser(user models.User) error
	FindUser(id uuid.UUID) (*models.User, error)
}
