package pg

import (
	"github.com/google/uuid"
	"gophermart/internal/domain/models"
)

func (p Pg) CreateWithdrawTransaction(tr models.Transaction) error {
	return nil
}

func (p Pg) FindUserWithdrawTransactions(user uuid.UUID) ([]models.Transaction, error) {
	return nil, nil
}

func (p Pg) FindUserWithdrawTransaction(user uuid.UUID, number string) (*models.Transaction, error) {
	return nil, nil
}
