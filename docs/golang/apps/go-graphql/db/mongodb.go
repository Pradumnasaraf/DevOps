package db

import (
	"context"
	"fmt"
	"go-grapgql/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var Ctx context.Context

// MongoDB is a struct that holds the MongoDB client
type MongoDB struct {
	Client *mongo.Client
}

// NewMongoDB creates a new MongoDB client and returns it
func NewMongoDB() *MongoDB {
	// Create a context
	Ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Collection = client.Database("test").Collection("videos")

	return &MongoDB{
		Client: client,
	}
}

func (db *MongoDB) Save(video *model.Video) { // db *MongoDB is a receiver. It's like a method of the struct. we can call it like db.Save()
	result, err := Collection.InsertOne(Ctx, video)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in DB with id:", result.InsertedID)

}

func (db *MongoDB) FindAll() (videos []*model.Video) {
	cursor, err := Collection.Find(Ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(Ctx)
	var result []*model.Video
	for cursor.Next(Ctx) {
		var video *model.Video
		err := cursor.Decode(&video)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, video)
	}

	return result
}
