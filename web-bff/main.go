package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yahya077/portfolio-microservice-app/web-bff/module/response"
	"github.com/yahya077/portfolio-microservice-app/web-bff/service/about"
	"os"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/about", func(ctx *fiber.Ctx) error {
		c := about.New().HandleGet(ctx)

		if c.HasError() {
			return ctx.SendStatus(c.GetCode())
		}

		return ctx.Status(c.GetCode()).JSON(response.NewResponse(response.WithSuccess(), response.WithData(c.GetSchema())))
	})

	_ = app.Listen(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
}
