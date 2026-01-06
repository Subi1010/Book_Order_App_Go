package services

import (
	"book_order_app/config"
	"book_order_app/middleware"
	"book_order_app/models"
	"errors"

	"gorm.io/gorm"
)

var userLogger = middleware.GetLogger()

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
		userLogger.WithError(err).WithField("username", req.Username).Error("Error creating user")
		return nil, errors.New("failed to create user")
	}

	userLogger.WithFields(map[string]interface{}{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
	}).Info("Successfully registered user")
	return &user, nil
}

// Login authenticates a user
func (us *userService) Login(req models.LoginRequest) (*models.User, error) {
	var user models.User
	if err := us.dbHandler.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userLogger.WithField("username", req.Username).Warn("Login attempt with invalid username")
			return nil, errors.New("invalid username or password")
		}
		userLogger.WithError(err).WithField("username", req.Username).Error("Error finding user during login")
		return nil, errors.New("failed to authenticate")
	}

	// Check password
	if err := user.CheckPassword(req.Password); err != nil {
		userLogger.WithField("username", req.Username).Warn("Login attempt with invalid password")
		return nil, errors.New("invalid username or password")
	}

	userLogger.WithFields(map[string]interface{}{
		"user_id":  user.ID,
		"username": user.Username,
	}).Info("User logged in successfully")
	return &user, nil
}

// GetByUsername retrieves a user by username
func (us *userService) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := us.dbHandler.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		userLogger.WithError(err).WithField("username", username).Error("Error finding user by username")
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
		userLogger.WithError(err).WithField("user_id", id).Error("Error finding user by ID")
		return nil, errors.New("failed to retrieve user")
	}
	return &user, nil
}
