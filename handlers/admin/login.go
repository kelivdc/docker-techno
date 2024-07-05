package admin

import (
	"github.com/gofiber/fiber/v2"
	"techno.com/db"
	"techno.com/lib"
	"techno.com/models"
)

func AdminLogin(c *fiber.Ctx) error {
	return c.Render("admin/login", fiber.Map{
		"Title": "Login",
	})
}

func AdminProses(c *fiber.Ctx) error {
	email := c.FormValue("email")
	var admin models.Admin
	db.DB.Db.Where("email = ?", email).First(&admin)
	if admin.Name == "" {
		return c.SendString("<span class='text-red-600'>Username or password not match</span>")
	}
	password := c.FormValue("password")
	if !lib.CheckPasswordHash(password, admin.Password) {
		return c.SendString("<span class='text-red-600'>Username or password not match</span>")
	}

	c.Response().Header.Add("HX-Redirect", "/admin/dashboard")
	return c.SendString("okeh")
}

func AdminCreate(c *fiber.Ctx) error {
	hash, _ := lib.HashPassword("12345678")
	admin := models.Admin{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: hash,
	}
	db.DB.Db.Create(&admin)
	return c.SendString("Sukses")
}
