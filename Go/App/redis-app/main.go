package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type User struct {
	UserName string `json:"username" validate:"required"`
}

// Load .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

// Redis client
func redisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return rdb

}

func main() {
	rdb := redisClient()

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API is running",
		})
	})

	// Root Route - For checking if the API is running
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API is running",
		})
	})

	// Get user by username Route
	app.Get("/get/:username", func(c *fiber.Ctx) error {

		params := c.Params("username")

		val, err := rdb.Get(ctx, params).Result()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}

		return c.JSON(fiber.Map{
			"message": "user found with name: " + val,
		})
	})

	// Create user Route
	app.Post("/post", func(c *fiber.Ctx) error {

		var user User

		err := c.BodyParser(&user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error parsing JSON. Make sure the request body is valid",
			})
		}
		if user.UserName == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Username key is required",
			})
		}

		err = rdb.Set(ctx, user.UserName, user.UserName, 3600*time.Second).Err()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error saving user, please try again",
			})
		}

		return c.JSON(fiber.Map{
			"message": "User created",
		})
	})

	// Delete user Route
	app.Delete("/delete/:username", func(c *fiber.Ctx) error {

		params := c.Params("username")

		err := rdb.Del(ctx, params).Err()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error deleting user, please try again",
			})
		}

		return c.JSON(fiber.Map{
			"message": "User deleted",
		})
	})

	app.Get("/getall", func(c *fiber.Ctx) error {

		val, err := rdb.Keys(ctx, "*").Result()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error getting users, please try again",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Users found",
			"users":   val,
		})
	})

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))

}
