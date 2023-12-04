package models

import (
	"fmt"
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

func GetOrderStatusByValue(status string) (OrderStatus, error) {
	switch status {
	case "NEW":
		return New, nil
	case "PROCESSING":
		return Processing, nil
	case "INVALID":
		return Invalid, nil
	case "PROCESSED":
		return Processed, nil
	default:
		return "", fmt.Errorf("status invalid")
	}
}
