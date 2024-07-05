package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"techno.com/handlers"
	admin_login "techno.com/handlers/admin"
)

func InitRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	app.Get("/", handlers.Home)
	admin := app.Group("admin")
	admin.Get("/", admin_login.AdminLogin)
	admin.Get("/dashboard", admin_login.AdminDashboard)

	admin.Get("/create-admin", admin_login.AdminCreate)
	admin.Post("/proses", admin_login.AdminProses)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("404 Not found!")
	})
}
