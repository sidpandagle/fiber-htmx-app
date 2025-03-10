package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "FlowBite",
		}, "layouts/main")
	})

	app.Get("/pricing", func(c *fiber.Ctx) error {
		return c.Render("pricing", fiber.Map{
			"Title": "Pricing",
		}, "layouts/main")
	})

	app.Get("/services", func(c *fiber.Ctx) error {
		return c.Render("services", fiber.Map{
			"Title": "Services",
		}, "layouts/main")
	})

	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.Render("contact", fiber.Map{
			"Title": "Contact",
		}, "layouts/main")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"Title": "About",
		}, "layouts/main")
	})

	app.Get("/report", func(c *fiber.Ctx) error {
		return c.Render("report", fiber.Map{
			"Title": "Report",
		}, "layouts/main")
	})

	log.Fatal(app.Listen(getPort()))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}
