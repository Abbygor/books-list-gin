package app

import (
	"github.com/Abbygor/books-list-gin/controllers"
)

var bookController controllers.BookController

func mapUrls() {

	router.GET("/books", bookController.GetBooks)
	router.GET("/books/:book_id", bookController.GetBook)
	router.POST("/books", bookController.AddBook)
	router.PUT("/books", bookController.UpdateBook)
	router.DELETE("books/:book_id", bookController.DeleteBook)
}
