package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed" default:"false"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("Hello, World!")

	DbConnection()
	Router()

}

func DbConnection() {
	MongoUri := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MongoUri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to DB")

	collection = client.Database("golang_db").Collection("todos")

}

func Router() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/api/todos", GetTodos)
	app.Post("/api/todos", CreateTodo)
	app.Patch("/api/todos/:id", UpdateTodo)
	app.Delete("/api/todos/:id", DeleteTodo)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	if err := app.Listen(":" + PORT); err != nil {
		log.Fatal(err)
	}
}

func GetTodos(c *fiber.Ctx) error {
	log.Println("Path: GetTodos entered")
	var todos []Todo
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cursor, context.Background())
	return c.Status(200).JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	log.Println("Path: CreateTodo entered")

	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		log.Fatal(err)
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Body is empty"})
	}

	one, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Cannot insert todo"})
	}

	todo.ID = one.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	log.Println("Path: UpdateTodo entered")

	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "ID is empty"})
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})

	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	col := collection.FindOneAndUpdate(context.Background(), filter, update)
	if col.Err() != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Todo updated"})
}

func DeleteTodo(c *fiber.Ctx) error {
	log.Println("Path: DeleteTodo entered")

	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "ID is empty"})
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Todo deleted"})
}
