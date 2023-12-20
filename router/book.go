package router

import "github.com/gofiber/fiber/v2"

func BookRouter(api fiber.Router) {
	route := api.Group("/book")

	route.Get("/")
}
