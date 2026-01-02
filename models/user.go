package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRole represents the role of a user
type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

// User represents a user in the system
type User struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index" format:"date-time"`
	Username  string     `json:"username" binding:"required" gorm:"uniqueIndex;not null" example:"johndoe"`
	Password  string     `json:"-" gorm:"not null"` // "-" means this field won't be included in JSON responses
	Role      UserRole   `json:"role" gorm:"type:varchar(20);not null;default:'user'" example:"user"`
}

// HashPassword hashes the user's password before saving
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares the provided password with the hashed password
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// BeforeCreate is a GORM hook that runs before creating a user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.HashPassword()
}

// LoginRequest represents the login request payload
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// RegisterRequest represents the registration request payload
type RegisterRequest struct {
	Username string   `json:"username" binding:"required" example:"johndoe"`
	Password string   `json:"password" binding:"required,min=6" example:"password123"`
	Role     UserRole `json:"role" binding:"required,oneof=admin user" example:"user"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  User   `json:"user"`
}
