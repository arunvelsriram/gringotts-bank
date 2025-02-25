package customer

import (
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	name string
	addr string
}

func (s Server) Run() error {
	app := fiber.New(fiber.Config{AppName: s.name})

	app.Use(otelfiber.Middleware())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"health": "ok"})
	})

	return app.Listen(s.addr)
}

func NewServer(name, addr string) Server {
	return Server{
		name: name,
		addr: addr,
	}
}
