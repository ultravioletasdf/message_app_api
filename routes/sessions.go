package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func setSession(c fiber.Ctx) error {
	return c.SendString("123")
}
func checkSession(c fiber.Ctx) error {
	session := c.Get("Authorization")
	fmt.Println(session)
	if session == "" {
		return c.SendString("0")
	} else {
		return c.SendString("1")
	}
}
func clearSession(c fiber.Ctx) error {
	return c.Status(200).Send(nil)
}
