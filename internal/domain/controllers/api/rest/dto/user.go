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
	UserId uuid.UUID
}

type GetOrders struct {
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    *int      `json:"accrual,omitempty"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type GetUserBalance struct {
	Current   float64 `json:"current"`
	Withdrawn float64 `json:"withdrawn"`
}

type UserBalanceWithdraw struct {
	UserId uuid.UUID
	Order  string  `json:"order"`
	Sum    float64 `json:"sum"`
}

type UserWithdraws struct {
	Order       string    `json:"order"`
	Sum         float64   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

type GetUser struct {
	Uid       uuid.UUID
	Login     string
	CreatedAt *time.Time
	DeletedAt *time.Time
}
