package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Date   string `json:"date"`
}

var books = []Book{}

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Book mm", Author: "BAS", Date: "2021-09-01"})
	books = append(books, Book{ID: 2, Title: "Book nn", Author: "BAS", Date: "2021-09-02"})
	books = append(books, Book{ID: 3, Title: "Book oo", Author: "BAS", Date: "2021-09-03"})

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", postBook)
	app.Put("/books/:id", PutBook)
	app.Delete("/books/:id", DeleteBook)

	app.Listen(":8080")

}
