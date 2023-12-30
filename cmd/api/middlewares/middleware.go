package middlewares

import(
	"github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/envvar"
)

func SetupMiddleware(app *fiber.App){
	app.Use("", NewEnvVar)
}