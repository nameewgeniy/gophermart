package dto

import (
	"github.com/google/uuid"
	"time"
)

type RegisterUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken  string
	RefreshToken string
}

type CreateOrder struct {
	Number string
}

type GetOrders struct {
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    int       `json:"accrual"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type GetUserBalance struct {
	Current   float64 `json:"current"`
	Withdrawn int     `json:"withdrawn"`
}

type UserBalanceWithdraw struct {
	Order string `json:"order"`
	Sum   int    `json:"sum"`
}

type UserWithdraws struct {
	Order       string    `json:"order"`
	Sum         int       `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

type GetUser struct {
	Uid       uuid.UUID
	Login     string
	CreatedAt *time.Time
	DeletedAt *time.Time
}
