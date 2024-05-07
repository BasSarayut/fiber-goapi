package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getPublishers(c *fiber.Ctx) error {

	return c.JSON(publishers)
}

func getPublisher(c *fiber.Ctx) error {
	publisherID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}

	for _, publisher := range publishers {
		if publisher.ID == publisherID {
			return c.JSON(publisher)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Publisher not found")

}

func postPublisher(c *fiber.Ctx) error {
	publisher := new(Publisher)

	if err := c.BodyParser(publisher); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	publisher.ID = len(publishers) + 1
	publishers = append(publishers, *publisher)

	return c.JSON(publisher)
}

func PutPublisher(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	publisherUpdate := new(Publisher)
	if err := c.BodyParser(publisherUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, publisher := range publishers {
		if publisher.ID == id {
			publishers[i].Name = publisherUpdate.Name
		}
	}

	return c.JSON(publishers)
}

func DeletePublisher(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for i, publisher := range publishers {
		if publisher.ID == id {
			publishers = append(publishers[:i], publishers[i+1:]...)
			return c.SendStatus(fiber.StatusOK)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
