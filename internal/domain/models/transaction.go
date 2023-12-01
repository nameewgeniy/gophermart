package models

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	Order       string
	Sum         money.Money
	ProcessedAt time.Time
}
