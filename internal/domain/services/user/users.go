package user

import (
	"fmt"
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
	UserGetOrders(userId uuid.UUID) ([]dto.GetOrders, error)
	UserBalance() (dto.GetUserBalance, error)
	UserBalanceWithdraw(bl dto.UserBalanceWithdraw) error
	UserWithdraws() ([]dto.UserWithdraws, error)
}

type User struct {
	s repositories.UserRepository
	o repositories.OrderRepository
	a auth.AuthService
}

func NewUserService(s repositories.UserRepository, o repositories.OrderRepository, a auth.AuthService) *User {
	return &User{
		s: s,
		o: o,
		a: a,
	}
}

func (u User) Register(us dto.RegisterUser) error {

	passwordHash, err := u.a.PasswordHash(us.Password)
	if err != nil {
		return err
	}

	return u.s.CreateUser(models.User{
		Login:        us.Login,
		PasswordHash: passwordHash,
	})
}

func (u User) Login(us dto.LoginUser) (dto.LoginUserResponse, error) {

	user, err := u.s.FindUserByLogin(us.Login)

	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	passwordHash, err := u.a.PasswordHash(us.Password)
	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	if user.PasswordHash != passwordHash {
		return dto.LoginUserResponse{}, fmt.Errorf("ошибка логина ли пароля")
	}

	tokens, err := u.a.TokenPair(*user)

	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	return dto.LoginUserResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (u User) UserCreateOrders(ord dto.CreateOrder) error {

	// Todo сначала найти существующий order

	exist, err := u.o.FindUserOrder(ord.UserId, ord.Number)

	if err != nil {
		return err
	}

	if exist != nil {
		return err // Вернуть ошибку, что такой order уже существует
	}

	om := models.Order{
		Id:         uuid.UUID{},
		UserId:     ord.UserId,
		Number:     ord.Number,
		Status:     models.New,
		UploadedAt: time.Now(),
	}

	if err = u.o.CreateOrder(om); err != nil {
		return err
	}

	return nil
}

func (u User) UserGetOrders(id uuid.UUID) ([]dto.GetOrders, error) {
	var res []dto.GetOrders

	ordrs, err := u.o.FindUserOrders(id)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(ordrs); i++ {
		res = append(res, dto.GetOrders{
			Number:     ordrs[i].Number,
			Status:     string(ordrs[i].Status),
			Accrual:    ordrs[i].Accrual,
			UploadedAt: ordrs[i].UploadedAt,
		})
	}

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
