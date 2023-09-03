package postgres

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Orders   []Order `gorm:"foreignKey:UserID"`
}

func (u *User) setPassword(raw string) error {
	hashed, err := hashPassword(raw)
	if err != nil {
		return err
	}

	u.Password = hashed
	return nil
}

func (u *User) checkPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func hashPassword(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	return string(hashed), err
}

type Status int

const (
	New Status = iota
	Processing
	Invalid
	Processed
)

func (s Status) String() string {
	return []string{"NEW", "PROCESSING", "INVALID", "PROCESSED"}[s]
}

type Order struct {
	gorm.Model
	ID      uint    `gorm:"primaryKey"`
	Number  string  `gorm:"unique"`
	Status  Status  `gorm:"default:0"`
	Accrual float64 `gorm:"default:0"`
	UserID  uint

	Transactions []Transaction `gorm:"foreignKey:OrderID"`
}

type Transaction struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	OrderID uint
	Amount  float64
}
