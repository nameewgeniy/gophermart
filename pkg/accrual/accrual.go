package accrual

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type AccrualClient interface {
	GetLoyaltyPoints(orderNumber string) (*AccrualResponse, error)
}

type AccrualClientImpl struct {
	baseAPIURL string
}

type AccrualResponse struct {
	Order      string  `json:"order,omitempty"`
	Status     string  `json:"status,omitempty"`
	Accrual    float64 `json:"accrual,omitempty"`
	RetryAfter time.Duration
}

func New(baseAPIURL string) *AccrualClientImpl {
	return &AccrualClientImpl{
		baseAPIURL: baseAPIURL,
	}
}

func (c AccrualClientImpl) GetLoyaltyPoints(orderNumber string) (*AccrualResponse, error) {

	url := c.baseAPIURL + orderNumber

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case 200:
		var accrualResponse AccrualResponse
		if err := json.NewDecoder(response.Body).Decode(&accrualResponse); err != nil {
			return nil, err
		}
		return &accrualResponse, nil

	case 204:
		return nil, ErrOrderNorRegistered

	case 429:
		retryAfter, _ := strconv.Atoi(response.Header.Get("Retry-After"))

		return &AccrualResponse{
			RetryAfter: time.Duration(retryAfter),
		}, ErrTooManyRequests

	case 500:
		return nil, fmt.Errorf("500 Internal Server Error")

	default:
		return nil, fmt.Errorf("unexpected response status: %s", response.Status)
	}
}
