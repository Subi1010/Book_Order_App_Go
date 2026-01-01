package controllers

import (
	"net/http"

	"book_order_app/models"
	"book_order_app/services"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service services.BookService
}

func InitializeBookController() *BookController {
	bookService := services.NewBookService()
	return &BookController{service: bookService}
}

// GetBooks godoc
// @Summary Get all books
// @Description Get a list of all books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func (bc *BookController) GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, bc.service.GetAll())
}

// AddBook godoc
// @Summary Add a new book
// @Description Create a new book with the provided information
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book information"
// @Success 201 {object} models.Book
// @Failure 400 {object} map[string]string
// @Router /books [post]
func (bc *BookController) AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := bc.service.Create(book)
	c.JSON(http.StatusCreated, created)
}
