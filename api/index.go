package handler

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	helloworld "io.github.com/fiber-company-profile/app/modules/hello-world"
)

// for vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	app := fiber.New()
	routes(app)
	return adaptor.FiberApp(app)
}

func routes(app *fiber.App) {
	app.Static("/", "../public")

	api := app.Group("/api")

	hello_world_api := api.Group("/hello-world")
	hello_world_api.Get("", helloworld.HelloWorld)

	app.Get("/config", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"SERVER_PORT": os.Getenv("SERVER_PORT"),
		})
	})
}
