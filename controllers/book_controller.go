package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Abbygor/books-list-gin/models"
	"github.com/Abbygor/books-list-gin/services"
	"github.com/Abbygor/books-list-gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookController struct{}

var bookService services.BookService

var validate *validator.Validate

func init() {
	validate = validator.New()
}

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

	apiErr := validateVariable(c.Param("book_id"))

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	book_id, _ := strconv.Atoi(c.Param("book_id"))

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

	apiErr := validateStruct(new_book)

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

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

	apiErr := validateStruct(update_book)

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	book, err := bookService.UpdateBook(update_book)

	if err != nil {
		utils.RespondError(c, err)
		return
	}
	utils.Respond(c, http.StatusOK, book)
}

func (bc BookController) DeleteBook(c *gin.Context) {
	log.Println("Controller DeleteBook")

	apiErr := validateVariable(c.Param("book_id"))

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	book_id, _ := strconv.Atoi(c.Param("book_id"))

	book, apiErr := bookService.DeleteBook(book_id)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, book)
}

func validateStruct(object interface{}) *utils.ApplicationError {

	err := validate.Struct(object)
	messages := []string{}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, err.Namespace()+" "+err.Tag()+" "+err.Kind().String()+" "+fmt.Sprintf("%v", err.Value()))
		}
		apiErr := &utils.ApplicationError{
			Message:    strings.Join(messages[:], ","),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		return apiErr
	}
	return nil
}

func validateVariable(variable interface{}) *utils.ApplicationError {
	errs := validate.Var(variable, "required,number")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		apiErr := &utils.ApplicationError{
			Message:    strings.Replace(errs.Error(), "''", "Book.ID", 2),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		return apiErr
	}
	return nil
}
