package services

import (
	"log"
	"net/http"

	"github.com/Abbygor/books-list-gin/models"
	"github.com/Abbygor/books-list-gin/repositories"
	"github.com/Abbygor/books-list-gin/utils"
)

type BookService struct{}

var bookRepository repositories.BookRepository

func (bs BookService) GetBooks() ([]models.Book, *utils.ApplicationError) {
	log.Println("Service GetBooks")

	return bookRepository.GetBooks()
}

func (bs BookService) GetBook(book_id int) (models.Book, *utils.ApplicationError) {
	log.Printf("Service GetBook %v", book_id)

	return bookRepository.GetBook(book_id)
}

func (bs BookService) AddBook(new_book models.Book) (models.Book, *utils.ApplicationError) {
	log.Println("Service AddBook")
	if new_book.Author == "" || new_book.Title == "" || new_book.Year == "" {
		return models.Book{}, &utils.ApplicationError{
			Message:    "Book error",
			StatusCode: http.StatusBadRequest,
			Code:       "book_error_content",
		}
	}
	return bookRepository.AddBook(new_book)
}

func (bs BookService) UpdateBook(update_book models.Book) (models.Book, *utils.ApplicationError) {
	log.Println("Service UpdateBook")
	if update_book.ID == 0 || update_book.Author == "" || update_book.Title == "" || update_book.Year == "" {
		return models.Book{}, &utils.ApplicationError{
			Message:    "All Book fields are required.",
			StatusCode: http.StatusBadRequest,
			Code:       "book_error_content",
		}
	}

	return bookRepository.UpdateBook(update_book)
}

func (bs BookService) DeleteBook(book_id int) (int, *utils.ApplicationError) {
	log.Printf("Service DeleteBook %v", book_id)

	return bookRepository.DeleteBook(book_id)

}
