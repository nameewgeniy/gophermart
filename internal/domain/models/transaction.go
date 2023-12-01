package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	Order       string
	Sum         uint
	ProcessedAt time.Time
}
