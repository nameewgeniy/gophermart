package user

import (
	"fmt"
	"github.com/Rhymond/go-money"
	_ "github.com/Rhymond/go-money"
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
	UserBalance(id uuid.UUID) (dto.GetUserBalance, error)
	UserBalanceWithdraw(bl dto.UserBalanceWithdraw) error
	UserWithdraws(id uuid.UUID) ([]dto.UserWithdraws, error)
}

type User struct {
	s repositories.UserRepository
	o repositories.OrderRepository
	w repositories.WithdrawTransactionRepository
	a auth.AuthService
}

func NewUserService(
	s repositories.UserRepository,
	o repositories.OrderRepository,
	w repositories.WithdrawTransactionRepository,
	a auth.AuthService,
) *User {
	return &User{
		s: s,
		o: o,
		w: w,
		a: a,
	}
}

// Register регистрация пользователя
func (u User) Register(us dto.RegisterUser) error {

	passwordHash, err := u.a.PasswordHash(us.Password)
	if err != nil {
		return err
	}

	err = u.s.CreateUser(models.User{
		Id:           uuid.New(),
		Login:        us.Login,
		PasswordHash: passwordHash,
	})

	if err != nil {
		return err // TODO отловить ошибку, о том, что пользователь уже существует
	}

	return nil
}

// Login аутентификация пользователя
func (u User) Login(us dto.LoginUser) (dto.LoginUserResponse, error) {

	user, err := u.s.FindUserByLogin(us.Login)

	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	if user == nil {
		return dto.LoginUserResponse{}, fmt.Errorf("user not found")
	}

	if !u.a.ValidatePassword(us.Password, user.PasswordHash) {
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

// UserCreateOrders загрузка пользователем номера заказа для расчёта;
func (u User) UserCreateOrders(ord dto.CreateOrder) error {

	// Todo сначала найти существующий order
	// При создании заказа, нужно по идентификатору заказа проверить бонусы для зачисления
	// если они есть, то пополнить балланс пользователя

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
		UploadedAt: time.Now().UTC(),
	}

	if err = u.o.CreateOrder(om); err != nil {
		return err
	}

	return nil
}

// UserGetOrders получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях;
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

// UserBalance получение текущего баланса счёта баллов лояльности пользователя;
func (u User) UserBalance(id uuid.UUID) (dto.GetUserBalance, error) {

	us, err := u.s.FindUser(id)

	if err != nil {
		return dto.GetUserBalance{}, err
	}

	Withdrawn := money.New(100, money.RUB)
	bl, err := us.Balance.Subtract(Withdrawn)

	if err != nil {
		return dto.GetUserBalance{}, err
	}

	return dto.GetUserBalance{
		Current:   bl.AsMajorUnits(),
		Withdrawn: Withdrawn.AsMajorUnits(),
	}, nil
}

// UserBalanceWithdraw запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа
func (u User) UserBalanceWithdraw(bl dto.UserBalanceWithdraw) error {

	exist, err := u.w.FindUserWithdrawTransaction(bl.UserId, bl.Order)
	if err != nil {
		return err
	}

	if exist != nil {
		return fmt.Errorf("такой заказ уже существует")
	}

	err = u.w.CreateWithdrawTransaction(models.Transaction{
		Id:          uuid.UUID{},
		UserId:      bl.UserId,
		Order:       bl.Order,
		Sum:         *money.NewFromFloat(bl.Sum, money.RUB),
		ProcessedAt: time.Now().UTC(),
	})

	if err != nil {
		return err
	}

	return nil
}

// UserWithdraws получение информации о выводе средств с накопительного счёта пользователем.
func (u User) UserWithdraws(id uuid.UUID) ([]dto.UserWithdraws, error) {

	var res []dto.UserWithdraws
	wd, err := u.w.FindUserWithdrawTransactions(id)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(wd); i++ {
		res = append(res, dto.UserWithdraws{
			Order:       wd[i].Order,
			Sum:         wd[i].Sum.AsMajorUnits(),
			ProcessedAt: wd[i].ProcessedAt,
		})
	}

	return res, nil
}
