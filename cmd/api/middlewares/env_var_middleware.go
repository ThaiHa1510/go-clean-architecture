package middlewars

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/envvar"
)
func NewEnvVar() fiber.Handler{
	return envvar.New(
    envvar.Config{
        ExportVars:  map[string]string{"testKey": "", "testDefaultKey": "testDefaultVal"},
        ExcludeVars: map[string]string{"excludeKey": ""},
    })
}