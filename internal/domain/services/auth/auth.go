package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gophermart/internal/domain/models"
	"gophermart/internal/domain/repositories"
	"sync"
	"time"
)

type AuthService interface {
	PasswordHash(pass string) (string, error)
	ValidatePassword(pass string, hash string) bool
	TokenPair(user models.User) (models.Token, error)
	VerifyToken(token string, tokenType models.TokenType) (*models.Claims, error)
	AccessToken(user models.User) (string, error)
	RefreshToken(user models.User) (string, error)
}

var (
	ErrUnexpectedMethod = errors.New("unexpected token signing method")
	ErrInvalidClaims    = errors.New("invalid token claims")
	ErrInvalidType      = errors.New("invalid token type")
	ErrInvalidToken     = errors.New("invalid token ")
)

var (
	Instance *JWTAuth
	once     sync.Once
)

type JWTAuth struct {
	ur         repositories.UserRepository
	secretKey  []byte
	accessTTL  int
	refreshTTL int
}

func NewAuthService(ur repositories.UserRepository, sk []byte, aTtl, rTtl int) *JWTAuth {
	once.Do(func() {
		Instance = &JWTAuth{
			ur:         ur,
			secretKey:  sk,
			accessTTL:  aTtl,
			refreshTTL: rTtl,
		}
	})

	return Instance
}

// PasswordHash generate hash from password
func (a JWTAuth) PasswordHash(pass string) (string, error) { //nolint
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// ValidatePassword check if password are equal
func (a JWTAuth) ValidatePassword(pass string, hash string) bool { //nolint
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))

	return err == nil
}

// TokenPair generate new token pair
func (a JWTAuth) TokenPair(user models.User) (models.Token, error) {
	accessToken, err := a.AccessToken(user)
	if err != nil {
		return models.Token{}, err
	}

	refreshToken, err := a.RefreshToken(user)
	if err != nil {
		return models.Token{}, err
	}

	return models.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// VerifyToken ...
func (a JWTAuth) VerifyToken(token string, tokenType models.TokenType) (*models.Claims, error) {
	parseen, err := jwt.ParseWithClaims(
		token,
		&models.Claims{},
		func(token *jwt.Token) (any, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, ErrUnexpectedMethod
			}

			return a.secretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := parseen.Claims.(*models.Claims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	if claims.TokenType != tokenType {
		return nil, ErrInvalidType
	}

	return claims, err
}

// AccessToken generate new access token
func (a JWTAuth) AccessToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(a.accessTTL) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		TokenType: models.TokenTypeAccess,
		UserId:    user.Id,
	})

	return token.SignedString(a.secretKey)
}

// RefreshToken generate new refresh token
func (a JWTAuth) RefreshToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(a.refreshTTL) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		TokenType: models.TokenTypeRefresh,
		UserId:    user.Id,
	})

	return token.SignedString(a.secretKey)
}
