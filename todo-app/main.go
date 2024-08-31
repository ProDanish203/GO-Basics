package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORt")

	app.Get("/api/v1/todos", getTodos)
	app.Post("/api/v1/todos", createTodo)
	app.Patch("/api/v1/todos/:id", updateTodo)
	app.Delete("/api/v1/todos/:id", deleteTodo)

	log.Fatal(app.Listen(":" + PORT))
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

// In memory code

// todos := []Todo{}

// app.Get("/api/v1/todos", func(c *fiber.Ctx) error {
// 	return c.Status(200).JSON(fiber.Map{"message": "Todos list", "data": todos})
// })
// app.Post("/api/v1/todos", func(c *fiber.Ctx) error {
// 	todo := &Todo{}

// 	if err := c.BodyParser(todo); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "error": err.Error()})
// 	}

// 	if todo.Body == "" {
// 		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "error": "Body is required"})
// 	}

// 	todo.ID = len(todos) + 1
// 	todos = append(todos, *todo)

// 	return c.Status(201).JSON(fiber.Map{"message": "Todo created", "data": todos})
// })
// app.Patch("/api/v1/todos/:id", func(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	todo := &Todo{}

// 	if err := c.BodyParser(todo); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "error": err.Error()})
// 	}

// 	for i, t := range todos {
// 		if fmt.Sprint(t.ID) == id {
// 			todos[i].Body = todo.Body
// 			todos[i].Completed = todo.Completed
// 			break
// 		}
// 	}

// 	return c.Status(200).JSON(fiber.Map{"message": "Todo updated", "data": todos})
// })
// app.Delete("/api/v1/todos/:id", func(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	for i, t := range todos {
// 		if fmt.Sprint(t.ID) == id {
// 			todos = append(todos[:i], todos[i+1:]...)
// 			break
// 		}
// 	}

// 	return c.Status(200).JSON(fiber.Map{"message": "Todo deleted", "data": todos})
// })
