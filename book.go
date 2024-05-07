package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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

func postBook(c *fiber.Ctx) error {
	book := new(Book)
	author := new(Author)
	publisher := new(Publisher)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	book.ID = len(books) + 1
	books = append(books, *book)
	author.ID = len(authors) + 1
	authors = append(authors, book.Author)
	publisher.ID = len(publishers) + 1
	publishers = append(publishers, book.Publisher)

	return c.JSON(book)
}

func PutBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == id {
			books[i].Title = bookUpdate.Title
			books[i].ID = bookUpdate.ID
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)

			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
