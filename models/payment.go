package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt  time.Time `json:"created_at"`
	CardNumber string    `json:"card_number"`
	ExpiryDate string    `json:"expiry_date"`
	CVV        string    `json:"cvv"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	UserID     uint      `json:"userId"` // Foreign key to User model

}
