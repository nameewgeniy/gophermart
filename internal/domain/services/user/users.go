package user

import (
	"github.com/google/uuid"
	"gophermart/internal/domain/controllers/api/rest/dto"
	"gophermart/internal/domain/models"
	"gophermart/internal/domain/repositories"
	"gophermart/internal/domain/services/auth"
	"time"
)

type UserService interface {
	Register(user dto.RegisterUser) error
	Login(user dto.LoginUser) (dto.LoginUserResponse, error)
	UserCreateOrders(ord dto.CreateOrder) error
	UserGetOrders() ([]dto.GetOrders, error)
	UserBalance() (dto.GetUserBalance, error)
	UserBalanceWithdraw(bl dto.UserBalanceWithdraw) error
	UserWithdraws() ([]dto.UserWithdraws, error)
}

type User struct {
	s repositories.UserRepository
	a auth.AuthService
}

func NewUserService(s repositories.UserRepository, a auth.AuthService) *User {
	return &User{
		s: s,
		a: a,
	}
}

func (u User) Register(us dto.RegisterUser) error {
	return u.s.CreateUser(models.User{
		Login:        us.Login,
		PasswordHash: us.Password,
	})
}

func (u User) Login(us dto.LoginUser) (dto.LoginUserResponse, error) {

	tokens, err := u.a.TokenPair(models.User{
		Id:           uuid.UUID{},
		Login:        "",
		PasswordHash: "",
	})

	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	return dto.LoginUserResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (u User) UserCreateOrders(ord dto.CreateOrder) error {
	return nil
}

func (u User) UserGetOrders() ([]dto.GetOrders, error) {
	var res []dto.GetOrders

	res = append(res, dto.GetOrders{
		Number:     "",
		Status:     "",
		Accrual:    0,
		UploadedAt: time.Time{},
	})

	return res, nil
}

func (u User) UserBalance() (dto.GetUserBalance, error) {
	return dto.GetUserBalance{
		Current:   0,
		Withdrawn: 0,
	}, nil
}

func (u User) UserBalanceWithdraw(bl dto.UserBalanceWithdraw) error {
	return nil
}

func (u User) UserWithdraws() ([]dto.UserWithdraws, error) {
	return []dto.UserWithdraws{}, nil
}
