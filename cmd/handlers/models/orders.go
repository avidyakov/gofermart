package models

import (
	"errors"
	"io"
	"regexp"
	"strconv"
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
	if !checkLuhn(string(o)) {
		return errors.New("invalid number according to the Luhn algorithm")
	}
	return nil
}

func doubleDigit(d int) int {
	dd := d * 2
	if dd > 9 {
		dd = dd - 9
	}
	return dd
}

func checkLuhn(s string) bool {
	n, _ := strconv.Atoi(s)
	sum := 0
	flip := false
	for n > 0 {
		if flip {
			sum += doubleDigit(n % 10)
		} else {
			sum += n % 10
		}
		n /= 10
		flip = !flip
	}
	return sum%10 == 0
}
