package pg

import (
	"github.com/google/uuid"
	"gophermart/internal/domain/models"
)

func (p Pg) CreateOrder(ord models.Order) error {
	return nil
}

func (p Pg) FindOrder(id uuid.UUID) (*models.Order, error) {
	return nil, nil
}

func (p Pg) FindUserOrder(user uuid.UUID, number string) (*models.Order, error) {
	return nil, nil
}

func (p Pg) FindUserOrders(user uuid.UUID) ([]models.Order, error) {
	return nil, nil
}
