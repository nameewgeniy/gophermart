package models

import (
	"github.com/Rhymond/go-money"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type TokenType string

const (
	TokenTypeAccess  TokenType = "accessToken"
	TokenTypeRefresh TokenType = "refreshToken"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

type Claims struct {
	jwt.StandardClaims
	TokenType TokenType
	UserId    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type User struct {
	Id           uuid.UUID
	Login        string
	PasswordHash string
	Balance      money.Money
	CreatedAt    time.Time
	DeletedAt    *time.Time
}
