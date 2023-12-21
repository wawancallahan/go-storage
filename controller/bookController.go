package controller

import (
	"pinjammodal/go-storage/dto"
	"pinjammodal/go-storage/service"
	"pinjammodal/go-storage/validation"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	service *service.BookService
}

func New(service *service.BookService) interface{} {
	return &BookController{service: service}
}

func (c *BookController) Create(ctx *fiber.Ctx) error {
	var book dto.Book

	if err := ctx.BodyParser(&book); err != nil {
		return err
	}

	errs := validation.SetupValidation(book)

	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  nil,
			"error": errs,
		})
	}

	response, err := c.service.Create(book)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}
