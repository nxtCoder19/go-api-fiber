package routes

import (
	"abc.io/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App){
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Put("/user/:userId", controllers.EditUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	app.Get("/user", controllers.GetAllUsers)
}