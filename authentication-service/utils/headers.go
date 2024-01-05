package utils

import "github.com/gofiber/fiber/v2"

func AddHeader(c *fiber.Ctx) {
	c.Set("Access-Control-Allow-Origin", c.Get("Origin"))
	c.Set("Access-Control-Allow-Credentials", "true")
}
