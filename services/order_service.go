package services

import (
	"book_order_app/config"
	"book_order_app/models"
	"log"
)

type OrderService interface {
	GetAll() []models.Order
	Create(order models.Order) models.Order
}

type orderService struct {
	dbHandler *config.DBHandler
}

func NewOrderService() OrderService {
	dbHandler := config.InitializeDBHandler()
	return &orderService{dbHandler: dbHandler}
}

func (os *orderService) GetAll() []models.Order {
	var orders []models.Order
	if err := os.dbHandler.DB.Find(&orders).Error; err != nil {
		log.Printf("Error fetching orders: %v", err)
		return []models.Order{}
	}
	return orders
}

func (os *orderService) Create(order models.Order) models.Order {
	if err := os.dbHandler.DB.Create(&order).Error; err != nil {
		log.Printf("Error creating order: %v", err)
		return order
	}
	return order
}
