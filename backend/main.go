package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server: Starting...")

	mode := os.Getenv("MODE")
	if mode == "dev" {
		fmt.Println("Running in dev mode")
		err := godotenv.Load(".env.dev")
		if err != nil {
			panic(err)
		}
		fmt.Println("Environment variables loaded")
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8000, http://0.0.0.0:8000, http://localhost:3000, http://0.0.0.0:3000",
	}))
	fmt.Println("Server: Created server")

	postgresUrl := os.Getenv("POSTGRES_URL")
	fmt.Printf("POSTGRES_URL:%v\n", postgresUrl)

	db, err := pgx.Connect(context.Background(), postgresUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())
	fmt.Println("Server: Connected to PostgreSQL")

	app.Get("/api", HelloWorld)
	app.Listen(":8000")
}

func HelloWorld(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}
