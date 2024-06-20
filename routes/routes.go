package routes

import "github.com/gofiber/fiber/v3"

func Add(app *fiber.App) {
	app.Get("/set_session", setSession)
	app.Get("/check_session", checkSession)
	app.Get("/clear_session", clearSession)
}
