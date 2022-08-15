package main

import (
	"github.com/go-chassis/go-chassis/v2"

	_ "github.com/go-chassis/go-chassis-extension/protocol/fiber4r"

	"github.com/go-chassis/openlog"
	"github.com/gofiber/fiber/v2"
)

func SayHello(c *fiber.Ctx) error {
	return c.SendString("hello. go chassis")
}

func main() {
	app := fiber.New()
	app.Get("/", SayHello)
	chassis.RegisterSchema("rest", app)
	if err := chassis.Init(); err != nil {
		openlog.Fatal("init failed." + err.Error())
		return
	}
	if err := chassis.Run(); err != nil {
		openlog.Fatal("run failed." + err.Error())
		return
	}
}
