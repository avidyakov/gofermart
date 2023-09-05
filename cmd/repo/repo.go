package repo

import (
	"errors"
	"time"
)

type Order struct {
	Number    string
	CreatedAt time.Time
	Status    string
	Accrual   float64
}

type Withdrawal struct {
	Order       string
	Sum         int
	ProcessedAt time.Time
}

type Repo interface {
	CreateUser(login, password string) (userID uint, err error)
	GetUser(login string) (userID uint, err error)
	CheckPassword(login, password string) (bool, error)

	CreateOrder(number string, userID uint) (orderID uint, err error)
	GetOrders(userID uint) ([]Order, error)

	MakeTransaction(orderNumber string, amount float64) error
	GetBalance(userID uint) (float64, error)
	GetUsed(userID uint) (float64, error)
	GetWithdrawals(userID uint) ([]Withdrawal, error)
}

var ErrUserExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user not found")
var ErrOrderExists = errors.New("order already exists")
var ErrOrderAlreadyUploaded = errors.New("order already uploaded by this user")
