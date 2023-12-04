package accrual

import (
	"errors"
)

var ErrOrderNorRegistered = errors.New("order not registered in the system")
var ErrTooManyRequests = errors.New("429 Too Many Requests")
