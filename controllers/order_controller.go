package controllers

import (
	"net/http"

	"book_order_app/models"
	"book_order_app/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
	bookService  services.BookService
}

func InitializeOrderController() *OrderController {
	orderService := services.NewOrderService()
	bookService := services.NewBookService()
	return &OrderController{
		orderService: orderService,
		bookService:  bookService,
	}
}

// GetOrders godoc
// @Summary Get all orders
// @Description Get a list of all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} models.Order
// @Router /orders [get]
func (oc *OrderController) GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, oc.orderService.GetAll())
}

// PlaceOrder godoc
// @Summary Place a new order
// @Description Create a new order for a book
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order information"
// @Success 201 {object} models.Order
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /orders [post]
func (oc *OrderController) PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !oc.bookService.Exists(order.BookID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	created := oc.orderService.Create(order)
	c.JSON(http.StatusCreated, created)
}
