package main

import (
	"log"

	"github.com/Leodf/leodf-go/internal/controller"
	"github.com/Leodf/leodf-go/internal/db"
	"github.com/bytedance/sonic"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	db, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Post("/clientes/:id/transacoes", controller.PostTransaction)
	app.Get("/clientes/:id/extrato", controller.GetStatment)
	app.Post("/reset-db", controller.ResetDB)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
