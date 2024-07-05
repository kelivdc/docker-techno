package admin

import (
	"github.com/gofiber/fiber/v2"
	"techno.com/constants"
)

func AdminDashboard(c *fiber.Ctx) error {
	return c.Render("admin/dashboard", fiber.Map{
		"Title": "Dashboard",
	}, constants.ADMIN_LAYOUT)
}
