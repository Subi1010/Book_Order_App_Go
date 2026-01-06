package services

import (
	"book_order_app/config"
	"book_order_app/middleware"
	"book_order_app/models"
)

var logger = middleware.GetLogger()

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
		logger.WithError(err).Error("Error fetching books")
		return []models.Book{}
	}
	logger.WithField("count", len(books)).Info("Successfully fetched books")
	return books
}

func (bs *bookService) Create(book models.Book) models.Book {
	if err := bs.dbHandler.DB.Create(&book).Error; err != nil {
		logger.WithError(err).WithField("book", book.Title).Error("Error creating book")
		return book
	}
	logger.WithFields(map[string]interface{}{
		"book_id": book.ID,
		"title":   book.Title,
	}).Info("Successfully created book")
	return book
}

func (bs *bookService) Exists(id uint) bool {
	var count int64
	bs.dbHandler.DB.Model(&models.Book{}).Where("id = ?", id).Count(&count)
	return count > 0
}
