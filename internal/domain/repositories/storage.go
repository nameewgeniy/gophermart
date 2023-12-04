package repositories

import (
	"context"
	"github.com/Rhymond/go-money"
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

type WithdrawTransactionRepository interface {
	CreateWithdrawTransaction(tr models.Transaction) error
	FindUserWithdrawTransaction(user uuid.UUID, number string) (*models.Transaction, error)
	FindUserWithdrawTransactions(user uuid.UUID) ([]models.Transaction, error)
	GetUserSumWithdraw(user uuid.UUID) (*money.Money, error)
}
