package handlers

import "github.com/gofiber/fiber/v2"

func HandleHealthCheck(c *fiber.Ctx) error {

	return c.Status(200).SendString("Server_Status : OK")

}
