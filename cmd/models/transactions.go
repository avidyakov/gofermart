package models

import (
	"encoding/json"
	"io"
	"time"
)

type Balance struct {
	Current   float64 `json:"current"`
	Withdrawn float64 `json:"withdrawn"`
}

type Withdrawal struct {
	Sum   float64 `json:"sum"`
	Order string  `json:"order"`
}

func (w *Withdrawal) Validate() error {
	return nil
}

func NewWithdrawal(raw io.Reader) (*Withdrawal, error) {
	var w Withdrawal
	decoder := json.NewDecoder(raw)

	decodeErr := decoder.Decode(&w)
	if decodeErr != nil {
		return nil, decodeErr
	}

	validationErr := w.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	return &w, nil
}

type WithdrawalOutput struct {
	Order       string    `json:"order"`
	Sum         int       `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}
