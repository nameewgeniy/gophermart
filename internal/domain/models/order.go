package models

import (
	"github.com/google/uuid"
	"time"
)

type OrderStatus string

const (
	New        OrderStatus = "NEW"
	Processing OrderStatus = "PROCESSING"
	Invalid    OrderStatus = "INVALID"
	Processed  OrderStatus = "PROCESSED"
)

type Order struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	Number     string
	Status     OrderStatus
	Accrual    *int
	UploadedAt time.Time
}
