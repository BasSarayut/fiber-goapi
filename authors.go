package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getAuthors(c *fiber.Ctx) error {

	return c.JSON(authors)

}

func getAuthor(c *fiber.Ctx) error {
	authorID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}

	for _, author := range authors {
		if author.ID == authorID {
			return c.JSON(author)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Author not found")

}

func postAuthor(c *fiber.Ctx) error {
	author := new(Author)

	if err := c.BodyParser(author); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	author.ID = len(authors) + 1
	authors = append(authors, *author)

	return c.JSON(author)
}

func PutAuthor(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	authorUpdate := new(Author)
	if err := c.BodyParser(authorUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, author := range authors {
		if author.ID == id {
			authors[i].Name = authorUpdate.Name
			authors[i].ID = authorUpdate.ID
			return c.JSON(author)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func DeleteAuthor(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for i, author := range authors {
		if author.ID == id {
			authors = append(authors[:i], authors[i+1:]...)

			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
