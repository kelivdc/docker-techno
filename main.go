package main

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"techno.com/config"
	"techno.com/db"
)

func main() {
	engine := html.New("./views", ".html")
	godotenv.Load(".env")
	app := fiber.New(fiber.Config{
		Views:     engine,
		Immutable: true,
	})
	// app.Static("/", "./public")
	app.Static("/", "./dist")
	config.InitRoutes(app)
	db.SetupDatabase()
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(cache.New(cache.Config{
		ExpirationGenerator: func(c *fiber.Ctx, cfg *cache.Config) time.Duration {
			newCacheTime, _ := strconv.Atoi(c.GetRespHeader("Cache-Time", "600"))
			return time.Second * time.Duration(newCacheTime)
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.Path())
		},
	}))

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	c.Response().Header.Add("Cache-Time", "6000")
	// 	return c.Render("index", fiber.Map{
	// 		"Title": "Jhony!",
	// 	}, "layouts/main")
	// })

	// app.Get("/about", func(c *fiber.Ctx) error {
	// 	c.Response().Header.Add("Cache-Time", "6000")
	// 	return c.Render("about", fiber.Map{
	// 		"Title": "Jhony!",
	// 	}, "layouts/main")
	// })

	// app.Get("/contact", func(c *fiber.Ctx) error {
	// 	c.Response().Header.Add("Cache-Time", "6000")
	// 	return c.Render("contact", fiber.Map{
	// 		"Title": "Contact",
	// 	}, "layouts/main")
	// })

	// app.Get("/login", func(c *fiber.Ctx) error {
	// 	c.Response().Header.Add("Cache-Time", "6000")
	// 	return c.Render("login", fiber.Map{
	// 		"Title": "Contact",
	// 	}, "layouts/main")
	// })

	// app.Post("/login", func(c *fiber.Ctx) error {
	// 	username := c.FormValue("username")
	// 	if username == "budy" {
	// 		c.Response().Header.Add("HX-Redirect", "/about")
	// 	}
	// 	return c.SendString("Username or password not match")
	// })

	// layout := "layouts/admin"
	// admin := app.Group("admin")
	// admin.Get("/login", handlers.AdminLogin)
	// admin.Post("/login", func(c *fiber.Ctx) error {
	// 	c.Response().Header.Add("HX-Redirect", "/admin/dashboard")
	// 	return c.Render("admin/dashboard", fiber.Map{
	// 		"Title": "Contact",
	// 	})
	// })
	// admin.Get("/dashboard", func(c *fiber.Ctx) error {
	// 	return c.Render("admin/dashboard", fiber.Map{
	// 		"Title": "Dashboard",
	// 	}, layout)
	// })
	// admin.Get("/category", func(c *fiber.Ctx) error {
	// 	return c.Render("admin/category", fiber.Map{
	// 		"Title": "Category",
	// 	}, layout)
	// })

	app.Listen(":3000")
}
