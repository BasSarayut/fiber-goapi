package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    Author    `json:"author"`
	Date      string    `json:"date"`
	Publisher Publisher `json:"publisher"`
}

type Author struct {
	ID   int    `json:"id_author"`
	Name string `json:"name"`
}

type Publisher struct {
	ID   int    `json:"id_publisher"`
	Name string `json:"name_publisher"`
}

var books = []Book{
	{ID: 1, Title: "The Hitchhiker's Guide to the Galaxy", Author: authors[0], Date: "1979", Publisher: publishers[0]},
	{ID: 2, Title: "The Lord of the Rings", Author: authors[1], Date: "1954", Publisher: publishers[1]},
	{ID: 3, Title: "The Hobbit", Author: authors[1], Date: "1937", Publisher: publishers[2]},
}

var authors = []Author{
	{ID: 1, Name: "Douglas Adams"},
	{ID: 2, Name: "J.R.R. Tolkien"},
	{ID: 3, Name: "George Allen & Unwin"},
}

var publishers = []Publisher{
	{ID: 1, Name: "Pan Books"},
	{ID: 2, Name: "George Allen & Unwin"},
	{ID: 3, Name: "HarperCollins"},
}

func main() {
	app := fiber.New()

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", postBook)
	app.Put("/books/:id", PutBook)
	app.Delete("/books/:id", DeleteBook)

	app.Get("/authors", getAuthors)
	app.Get("/authors/:id", getAuthor)
	app.Post("/authors", postAuthor)
	app.Put("/authors/:id", PutAuthor)
	app.Delete("/authors/:id", DeleteAuthor)

	app.Get("/publishers", getPublishers)
	app.Get("/publishers/:id", getPublisher)
	app.Post("/publishers", postPublisher)
	app.Put("/publishers/:id", PutPublisher)
	app.Delete("/publishers/:id", DeletePublisher)

	app.Listen(":8080")

}
