package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/mateigraura/wirebo-api/data/db"
	"github.com/mateigraura/wirebo-api/utils"
)

func main() {
	env := os.Args[1]

	app := fiber.New()

	app.Get("/", hello)

	utils.LoadEnvFile(env)

	conn := db.Connection()
	db.CreateSchema(conn)

	if err := app.Listen(utils.GetEnvFile()[utils.Port]); err != nil {
		log.Fatal(err)
	}
}

func hello(ctx *fiber.Ctx) {
	ctx.Send("Hello")
}
