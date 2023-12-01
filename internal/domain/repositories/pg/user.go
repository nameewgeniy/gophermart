package pg

import (
	"github.com/google/uuid"
	"gophermart/internal/domain/models"
)

func (p Pg) CreateUser(user models.User) error {
	return nil
}

func (p Pg) FindUser(id uuid.UUID) (*models.User, error) {
	return nil, nil
}

func (p Pg) FindUserByLogin(login string) (*models.User, error) {
	return nil, nil
}
