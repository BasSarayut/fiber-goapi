package main

import (
	"strconv"

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
	app.Get("/books/:Title", getAuthor)

	app.Listen(":8080")

}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}

	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Book not found")
}

func getAuthor(c *fiber.Ctx) error {
	author := c.Params("Author")

	for _, book := range books {
		if book.Author == author {
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Book not found")
}
