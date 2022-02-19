package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Abbygor/books-list-gin/models"
	"github.com/Abbygor/books-list-gin/services"
	"github.com/Abbygor/books-list-gin/utils"
	"github.com/gin-gonic/gin"
)

type BookController struct{}

var bookService services.BookService

func (bc BookController) GetBooks(c *gin.Context) {
	log.Println("Controller GetBooks")

	books, err := bookService.GetBooks()

	if err != nil {
		utils.RespondError(c, err)
		return
	}
	utils.Respond(c, http.StatusOK, books)
}

func (bc BookController) GetBook(c *gin.Context) {
	log.Println("Controller GetBook")

	book_id, err := strconv.Atoi(c.Param("book_id"))

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	book, apiErr := bookService.GetBook(book_id)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, book)
}

func (bc BookController) AddBook(c *gin.Context) {
	log.Println("Controller AddBook")
	new_book := models.Book{}
	c.Bind(&new_book)

	book, err := bookService.AddBook(new_book)

	if err != nil {
		utils.RespondError(c, err)
		return
	}
	utils.Respond(c, http.StatusOK, book)
}

func (bc BookController) UpdateBook(c *gin.Context) {
	log.Println("Controller UpdateBook")
	update_book := models.Book{}
	c.Bind(&update_book)

	book, err := bookService.UpdateBook(update_book)

	if err != nil {
		utils.RespondError(c, err)
		return
	}
	utils.Respond(c, http.StatusOK, book)
}

func (bc BookController) DeleteBook(c *gin.Context) {
	log.Println("Controller DeleteBook")

	book_id, err := strconv.Atoi(c.Param("book_id"))

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	book, apiErr := bookService.DeleteBook(book_id)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, book)
}
