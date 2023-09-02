package models

import "time"

type Order struct {
	Number    string    `json:"number"`
	CreatedAt time.Time `json:"uploaded_at"`
	Status    string    `json:"status"`
	Accrual   float64   `json:"accrual"`
}
