package models

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID
	Login        string
	PasswordHash string
}
