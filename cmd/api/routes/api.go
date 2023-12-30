package routes

import(
	"context"
)

func LoadApiRoutes(app *fiber.App){
	app.Route("/api", func(api fiber.Router){
		api.Get("/", controller.HomeIndex)
		api.Post("/user", controller.RegisterUserHandler)
	})
}