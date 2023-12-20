package controller

import (
	"pinjammodal/go-storage/dto"
	"pinjammodal/go-storage/service"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	service *service.BookService
}

var (
	validate = make(validation)
)

func New(service *service.BookService) interface{} {
	return &BookController{service: service}
}

func (c *BookController) Create(ctx *fiber.Ctx) error {
	var book dto.Book

	if err := ctx.BodyParser(&book); err != nil {
		return err
	}

	response, err := c.service.Create(book)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}
