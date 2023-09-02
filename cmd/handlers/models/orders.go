package models

import (
	"errors"
	"io"
	"regexp"
	"time"
)

type OrderOutput struct {
	Number    string    `json:"number"`
	CreatedAt time.Time `json:"uploaded_at"`
	Status    string    `json:"status"`
	Accrual   float64   `json:"accrual"`
}

type OrderInput string

func NewOrderInput(rawOrder io.Reader) (*OrderInput, error) {
	order, _ := io.ReadAll(rawOrder)
	orderInput := OrderInput(order)
	validationErr := orderInput.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	return &orderInput, nil
}

func (o OrderInput) Validate() error {
	matched, err := regexp.MatchString(`^[0-9]+$`, string(o))
	if err != nil || !matched {
		return errors.New("number must contain only integers")
	}

	return nil
}
