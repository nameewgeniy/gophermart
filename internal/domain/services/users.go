package services

import (
	"gophermart/internal/domain/controllers/api/rest/dto"
	"gophermart/internal/domain/models"
	"gophermart/internal/domain/repositories"
)

type UserService interface {
	Register(user dto.RegisterUser) error
	Login(user dto.LoginUser) error
	UserCreateOrders(ord dto.CreateOrder) error
	UserGetOrders() (dto.GetOrders, error)
	UserBalance() (dto.GetUserBalance, error)
	UserBalanceWithdraw(bl dto.UserBalanceWithdraw) error
	UserWithdraws() (dto.UserWithdraws, error)
}

type User struct {
	s repositories.UserRepository
}

func NewUserService(s repositories.UserRepository) *User {
	return &User{
		s: s,
	}
}

func (u User) Register(us dto.RegisterUser) error {
	return u.s.CreateUser(models.User{
		Login:        us.Login,
		PasswordHash: us.Password,
	})
}

func (u User) Login(us dto.LoginUser) error {
	return u.s.CreateUser(models.User{
		Login:        us.Login,
		PasswordHash: us.Password,
	})
}

func (u User) UserCreateOrders(ord dto.CreateOrder) error {
	return nil
}

func (u User) UserGetOrders() (dto.GetOrders, error) {
	return dto.GetOrders{}, nil
}

func (u User) UserBalance() (dto.GetUserBalance, error) {
	return dto.GetUserBalance{}, nil
}

func (u User) UserBalanceWithdraw(bl dto.UserBalanceWithdraw) error {
	return nil
}

func (u User) UserWithdraws() (dto.UserWithdraws, error) {
	return dto.UserWithdraws{}, nil
}
