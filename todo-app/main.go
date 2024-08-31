package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/v1/todos", getTodos)
	app.Post("/api/v1/todos", createTodo)
	app.Patch("/api/v1/todos/:id", updateTodo)
	app.Delete("/api/v1/todos/:id", deleteTodo)

	log.Fatal(app.Listen(":8000"))
}

func getTodos(c *fiber.Ctx) error {
	return c.SendString("Get Todos")
}

func createTodo(c *fiber.Ctx) error {
	return c.SendString("Create Todo")
}

func updateTodo(c *fiber.Ctx) error {
	// id := c.Params("id")
	return c.SendString("Update Todo")
}

func deleteTodo(c *fiber.Ctx) error {
	// id := c.Params("id")
	return c.SendString("Delete Todo")
}
