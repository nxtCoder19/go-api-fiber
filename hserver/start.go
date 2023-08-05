package hserver

import (
	"log"

	"abc.io/configs"
	"abc.io/routes"
	"github.com/gofiber/fiber/v2"
)

func Start() {
    app := fiber.New()

    // app.Get("/", func (c *fiber.Ctx) error {
    //     // return c.SendString("Hello, World!")
    //     return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    // })
    configs.ConnectDB()
    routes.UserRoute(app)

    log.Fatal(app.Listen(":3000"))
}