package repositories

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abbygor/books-list-gin/config"
	"github.com/Abbygor/books-list-gin/models"
	"github.com/Abbygor/books-list-gin/utils"
	"github.com/subosito/gotenv"
)

type BookRepository struct{}

func init() {
	gotenv.Load()
}

func (br BookRepository) GetBooks() ([]models.Book, *utils.ApplicationError) {
	log.Println("Repository GetBooks")

	db := config.ConnectDB()
	books := []models.Book{}
	book := models.Book{}

	rows, err := db.Query("select * from books")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	if err != nil {
		return []models.Book{}, &utils.ApplicationError{
			Message:    "Server Error",
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_server_error",
		}
	}

	return books, nil
}

func (br BookRepository) GetBook(book_id int) (models.Book, *utils.ApplicationError) {
	log.Printf("Repository GetBook %v", book_id)

	db := config.ConnectDB()

	book := models.Book{}

	row := db.QueryRow("select * from books where id = $1", book_id)

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	if err != nil || book.ID != book_id {
		return book, &utils.ApplicationError{
			Message:    fmt.Sprintf("Book %v not found", book_id),
			StatusCode: http.StatusNotFound,
			Code:       "book_not_found",
		}
	}

	return book, nil
}

func (br BookRepository) AddBook(new_book models.Book) (models.Book, *utils.ApplicationError) {
	log.Println("Repository AddBook")

	db := config.ConnectDB()

	err := db.QueryRow("insert into books(title, author, year) values($1, $2, $3) RETURNING id;",
		&new_book.Title, &new_book.Author, &new_book.Year).Scan(&new_book.ID)
	if err != nil {
		return models.Book{}, &utils.ApplicationError{
			Message:    "Book Insert fail",
			StatusCode: http.StatusInternalServerError,
			Code:       "server_error",
		}
	}

	return new_book, nil

}

func (br BookRepository) UpdateBook(update_book models.Book) (models.Book, *utils.ApplicationError) {
	log.Println("Repository UpdateBook")

	db := config.ConnectDB()

	result, err := db.Exec("update books set title = $1, author = $2, year = $3 where id = $4 RETURNING id;",
		&update_book.Title, &update_book.Author, &update_book.Year, &update_book.ID)

	if err != nil {
		return models.Book{}, &utils.ApplicationError{
			Message:    "Book Update fail",
			StatusCode: http.StatusBadRequest,
			Code:       "book_error_content",
		}
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil || rowsUpdated == 0 {
		return models.Book{}, &utils.ApplicationError{
			Message:    "Book Update 0 rows affected",
			StatusCode: http.StatusBadRequest,
			Code:       "book_error_content",
		}
	}

	return update_book, nil
}

func (br BookRepository) DeleteBook(book_id int) (int, *utils.ApplicationError) {
	log.Println("Repository DeleteBook")

	db := config.ConnectDB()

	result, err := db.Exec("delete from books where id = $1;", book_id)

	if err != nil {
		return 0, &utils.ApplicationError{
			Message:    "Book Delete fail",
			StatusCode: http.StatusBadRequest,
			Code:       "book_error_content",
		}
	}

	rowsDel, err := result.RowsAffected()

	if err != nil || rowsDel == 0 {
		return 0, &utils.ApplicationError{
			Message:    "Book Update 0 rows affected",
			StatusCode: http.StatusBadRequest,
			Code:       "book_error_content",
		}
	}

	return int(rowsDel), nil
}
