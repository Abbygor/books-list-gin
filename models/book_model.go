package models

type Book struct {
	ID     int    `json:"id" validate:"required,numeric"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Year   string `json:"year" validate:"required,numeric"`
}
