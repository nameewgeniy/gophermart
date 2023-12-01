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
	FindUserByLogin(login string) (*models.User, error)
}

type OrderRepository interface {
	CreateOrder(ord models.Order) error
	FindOrder(id uuid.UUID) (*models.Order, error)
	FindUserOrder(user uuid.UUID, number string) (*models.Order, error)
	FindUserOrders(user uuid.UUID) ([]models.Order, error)
}

type TransactionRepository interface {
	CreateTransaction(tr models.Transaction) error
	FindUserTransactions(user uuid.UUID) ([]models.Transaction, error)
}
