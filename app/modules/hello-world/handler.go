package helloworld

import (
	"github.com/gofiber/fiber/v2"
	"io.github.com/fiber-company-profile/app/dto/response"
)

func HelloWorld(c *fiber.Ctx) error {
	res := &response.Response{}
	res.Ok(c.Path(), "hello world!")

	return c.Status(res.Status).JSON(res)
}
