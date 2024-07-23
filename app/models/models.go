package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	UserID      uuid.UUID     `gorm:"column:user_id;type:uuid;primary_key;" json:"user_id"`
	FirstName   string        `gorm:"column:first_name" json:"first_name" validate:"required,alpha"`
	LastName    string        `gorm:"column:last_name" json:"last_name" validate:"required,alpha"`
	PhoneNumber string        `gorm:"column:phone_number;unique;not null" json:"phone_number" validate:"required,numeric,min=9,max=15"`
	Address     string        `gorm:"column:address" json:"address" validate:"required"`
	PIN         string        `gorm:"column:pin" json:"pin" validate:"required,numeric,min=6,max=6"`
	Balance     int           `gorm:"column:balance" json:"balance"`
	CreatedAt   time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *sql.NullTime `gorm:"colmun:updated_at" json:"updated_at"`
}

func (Users) TableName() string {
	return "users"
}

type Transactions struct {
	TransactionID   uuid.UUID     `gorm:"column:transaction_id;type:uuid;primary_key;" json:"transaction_id"`
	TopUpID         *uuid.UUID    `gorm:"column:top_up_id;type:uuid;" json:"top_up_id"`
	PaymentID       *uuid.UUID    `gorm:"column:payment_id;type:uuid;" json:"payment_id"`
	TransferID      *uuid.UUID    `gorm:"column:transfer_id;type:uuid;" json:"transfer_id"`
	UserID          uuid.UUID     `gorm:"column:user_id;type:uuid;" json:"user_id"`
	Amount          int           `gorm:"column:amount" json:"amount"`
	BalanceBefore   int           `gorm:"column:balance_before" json:"balance_before"`
	BalanceAfter    int           `gorm:"column:balance_after" json:"balance_after"`
	Remarks         string        `gorm:"column:remarks" json:"remarks"`
	TransactionType string        `gorm:"column:transaction_type" json:"transaction_type"`
	CreatedAt       time.Time     `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       *sql.NullTime `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at`
}

func (Transactions) TableName() string {
	return "transactions"
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"  validate:"required,numeric,min=9,max=15"`
	PIN         string `json:"pin" validate:"required,numeric,min=6,max=6"`
}

type TopUpRequest struct {
	Amount int `json:"amount" validate:"required,numeric"`
}

type PaymentRequest struct {
	Amount  int    `json:"amount" validate:"required,numeric"`
	Remarks string `json:"remarks"`
}

type TransferRequest struct {
	TargetUser string `json:"target_user" validate:"required"`
	Amount     int    `json:"amount" validate:"required,numeric"`
	Remarks    string `json:"remarks"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" validate:"required,alpha"`
	LastName  string `json:"last_name" validate:"required,alpha"`
	Address   string `json:"address"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Users{}, &Transactions{})
}
