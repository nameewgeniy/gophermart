package pg

import (
	"github.com/google/uuid"
	"gophermart/internal/domain/models"
)

func (p Pg) CreateTransaction(tr models.Transaction) error {
	return nil
}

func (p Pg) FindUserTransactions(user uuid.UUID) ([]models.Transaction, error) {
	return nil, nil
}
