package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	helloworld "io.github.com/fiber-company-profile/app/modules/hello-world"
)

// for vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	initEnvironment()

	handler().ServeHTTP(w, r)
}

// for local
func main() {
	app := fiber.New()

	initEnvironment()

	routes(app)

	log.Fatal(app.Listen(":" + os.Getenv("SERVER_PORT")))
}

func initEnvironment() {
	err := LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func handler() http.HandlerFunc {
	app := fiber.New()
	routes(app)
	return adaptor.FiberApp(app)
}

func routes(app *fiber.App) {
	app.Static("/", "../ui/dist")

	api := app.Group("/api")

	hello_world_api := api.Group("/hello-world")
	hello_world_api.Get("", helloworld.HelloWorld)
}

func LoadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set to System Environment so os.Getenv works
		os.Setenv(key, value)
	}

	return scanner.Err()
}
