package services

import (
	"book_order_app/config"
	"book_order_app/models"
	"errors"
	"log"

	"gorm.io/gorm"
)

type UserService interface {
	Register(req models.RegisterRequest) (*models.User, error)
	Login(req models.LoginRequest) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
}

type userService struct {
	dbHandler *config.DBHandler
}

func NewUserService() UserService {
	dbHandler := config.InitializeDBHandler()
	return &userService{dbHandler: dbHandler}
}

// Register creates a new user
func (us *userService) Register(req models.RegisterRequest) (*models.User, error) {
	// Check if user already exists
	var existingUser models.User
	if err := us.dbHandler.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("username already exists")
	}

	// Create new user
	user := models.User{
		Username: req.Username,
		Password: req.Password, // Will be hashed by BeforeCreate hook
		Role:     req.Role,
	}

	if err := us.dbHandler.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, errors.New("failed to create user")
	}

	return &user, nil
}

// Login authenticates a user
func (us *userService) Login(req models.LoginRequest) (*models.User, error) {
	var user models.User
	if err := us.dbHandler.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid username or password")
		}
		log.Printf("Error finding user: %v", err)
		return nil, errors.New("failed to authenticate")
	}

	// Check password
	if err := user.CheckPassword(req.Password); err != nil {
		return nil, errors.New("invalid username or password")
	}

	return &user, nil
}

// GetByUsername retrieves a user by username
func (us *userService) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := us.dbHandler.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		log.Printf("Error finding user: %v", err)
		return nil, errors.New("failed to retrieve user")
	}
	return &user, nil
}

// GetByID retrieves a user by ID
func (us *userService) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := us.dbHandler.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		log.Printf("Error finding user: %v", err)
		return nil, errors.New("failed to retrieve user")
	}
	return &user, nil
}
