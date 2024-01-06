package main

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/metrics")
	})

	app.Get("/metrics", func(c *fiber.Ctx) error {
		randomNumber := rand.Float64() // Generate a random number between 0 and 1

		// Prepare the metric with a comment
		metricComment := "# HELP random_number A random number between 0 and 1\n"
		metricValue := fmt.Sprintf("random_number %f\n", randomNumber)
		metrics := metricComment + metricValue

		return c.SendString(metrics)
	})

	// Start the server on port 8080
	app.Listen(":8080")
}
