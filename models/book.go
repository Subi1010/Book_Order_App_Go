package models

import (
	"time"
)

// Book represents a book in the system
type Book struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index" format:"date-time"`
	Title     string     `json:"title" binding:"required" gorm:"not null" example:"The Go Programming Language"`
	Author    string     `json:"author" binding:"required" gorm:"not null" example:"Alan A. A. Donovan"`
	Price     float64    `json:"price" binding:"required" gorm:"not null" example:"29.99"`
}

// CreateBookRequest represents the request body for creating a book
type CreateBookRequest struct {
	Title  string  `json:"title" binding:"required" example:"The Go Programming Language"`
	Author string  `json:"author" binding:"required" example:"Alan A. A. Donovan"`
	Price  float64 `json:"price" binding:"required" example:"29.99"`
}
