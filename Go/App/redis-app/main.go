package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type User struct {
	UserName string `json:"username" validate:"required"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func redisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return rdb

}

func main() {
	redisClient()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API is running",
		})
	})

	app.Get("/get/:username", func(c *fiber.Ctx) error {
		rdb := redisClient()

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
				"message": "name key is required",
			})
		}

		rdb := redisClient()

		err = rdb.Set(ctx, user.UserName, user.UserName, 60*time.Second).Err()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error saving user, please try again",
			})
		}

		return c.JSON(fiber.Map{
			"message": "User created",
		})
	})

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))

}
