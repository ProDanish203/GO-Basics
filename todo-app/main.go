package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Body      string             `json:"body"`
	Completed bool               `json:"completed"`
}

type UpdateTodoRequest struct {
	Completed *bool  `json:"completed,omitempty"`
	Body      string `json:"body,omitempty"`
}

var collection *mongo.Collection

func main() {
	// Create a server
	app := fiber.New()

	// Load Environment Variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	MONGO_URI := os.Getenv("MONGO_URI")

	// Confgure CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect((context.Background()), clientOptions)
	if err != nil {
		log.Fatal("Error connection to the Database: ", err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging the Database: ", err)
	}

	// Setup collection
	collection = client.Database("go-todo-app").Collection("todos")

	// Setup Routes
	app.Get("/api/v1/todos", getTodos)
	app.Post("/api/v1/todos", createTodo)
	app.Patch("/api/v1/todos/:id", updateTodo)
	app.Delete("/api/v1/todos/:id", deleteTodo)

	// Start Server
	log.Fatal(app.Listen(":" + PORT))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal("Error fetching todos: ", err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			log.Fatal("Error decoding todo: ", err)
		}
		todos = append(todos, todo)
	}

	return c.Status(200).JSON(fiber.Map{"message": "Todos list", "data": todos})
}

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "error": err.Error()})
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "error": "Body is required"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal("Error creating todo: ", err)
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(fiber.Map{"message": "Todo created", "data": todo})
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "error": "Invalid Todo ID"})
	}

	var body UpdateTodoRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "error": "Invalid request body"})
	}

	update := bson.M{}
	if body.Completed != nil {
		update["completed"] = *body.Completed
	}
	if body.Body != "" {
		update["body"] = body.Body
	}

	if len(update) == 0 {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "error": "At least one field (completed or body) must be provided"})
	}

	_, err = collection.UpdateByID(context.Background(), objectId, bson.M{"$set": update})
	if err != nil {
		log.Fatal("Error updating todo: ", err)
	}

	return c.Status(200).JSON(fiber.Map{"message": "Todo updated"})
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request", "error": "Invalid Todo ID"})
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectId})

	if err != nil {
		log.Fatal("Error deleting todo: ", err)
	}

	return c.Status(200).JSON(fiber.Map{"message": "Todo deleted"})
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
