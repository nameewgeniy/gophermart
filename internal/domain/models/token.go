package models

import (
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
