package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"os"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/get-contact", func(ctx *fiber.Ctx) error {
		result := make(map[string]interface{})

		jsonFile, err := os.Open("./data/contact.json")
		data, _ := io.ReadAll(jsonFile)
		defer jsonFile.Close()

		if err != nil {
			panic(err)
		}

		_ = json.Unmarshal(data, &result)

		return ctx.JSON(result)
	})

	_ = app.Listen(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
}
