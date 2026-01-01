package services

import (
	"book_order_app/config"
	"book_order_app/models"
	"log"
)

type BookService interface {
	GetAll() []models.Book
	Create(book models.Book) models.Book
	Exists(id uint) bool
}

type bookService struct {
	dbHandler *config.DBHandler
}

func NewBookService() BookService {
	dbHandler := config.InitializeDBHandler()
	return &bookService{dbHandler: dbHandler}
}

func (bs *bookService) GetAll() []models.Book {
	var books []models.Book
	if err := bs.dbHandler.DB.Find(&books).Error; err != nil {
		log.Printf("Error fetching books: %v", err)
		return []models.Book{}
	}
	return books
}

func (bs *bookService) Create(book models.Book) models.Book {
	if err := bs.dbHandler.DB.Create(&book).Error; err != nil {
		log.Printf("Error creating book: %v", err)
		return book
	}
	return book
}

func (bs *bookService) Exists(id uint) bool {
	var count int64
	bs.dbHandler.DB.Model(&models.Book{}).Where("id = ?", id).Count(&count)
	return count > 0
}
