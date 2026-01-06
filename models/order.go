package models

import (
	"time"
)

// Order represents an order in the system
type Order struct {
	ID           uint       `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" gorm:"index" format:"date-time"`
	BookID       uint       `json:"book_id" binding:"required" gorm:"not null" example:"1"`
	CustomerName string     `json:"customer_name" binding:"required" gorm:"not null" example:"John Doe"`
	Quantity     int        `json:"quantity" binding:"required" gorm:"not null" example:"2"`
}

// CreateOrderRequest represents the request body for creating an order
type CreateOrderRequest struct {
	BookID       uint   `json:"book_id" binding:"required" example:"1"`
	CustomerName string `json:"customer_name" binding:"required" example:"John Doe"`
	Quantity     int    `json:"quantity" binding:"required,min=1" example:"2"`
}
