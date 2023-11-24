package dto

import (
	"github.com/google/uuid"
	"time"
)

type RegisterUser struct {
	Login    string
	Password string
}

type LoginUser struct {
	RegisterUser
}

type CreateOrder struct {
	Id string
}

type GetOrders struct {
	Id string
}

type GetUserBalance struct {
}

type UserBalanceWithdraw struct {
}

type UserWithdraws struct {
}

type GetUser struct {
	Uid       uuid.UUID
	Login     string
	CreatedAt *time.Time
	DeletedAt *time.Time
}
