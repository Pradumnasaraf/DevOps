package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/pradumnasaraf/mongo-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURL = "mongodb://localhost:27017"
const databaseName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

//Connecting with MongoDB

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	checkNilErr(err)

	fmt.Println("MongoDB Connected sucessfully")

	collection = client.Database(databaseName).Collection(collectionName)
	fmt.Println("Collection instance is ready")

}

// MONGODB HELPERS

func insertOneMovie(movie model.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie)

	checkNilErr(err)
	fmt.Println("Inserted 1 movie in DB with id:", result.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	checkNilErr(err)
	fmt.Println("Modified count:", result.UpsertedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)

	checkNilErr(err)
	fmt.Println("Modified count:", result.DeletedCount)
}

func deleteAllMovies() int64 {
	filter := bson.D{{}}
	result, err := collection.DeleteMany(context.Background(), filter)

	checkNilErr(err)
	fmt.Println("Modified count:", result.DeletedCount)
	return result.DeletedCount
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteMyAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deltedCount := deleteAllMovies()
	json.NewEncoder(w).Encode(deltedCount)
}
func ServeHomepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is working fine 🚀."))
}
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	checkNilErr(err)
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		checkNilErr(err)

		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies

}
func GetMyAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func getOneMovies(movieId string) model.Netflix {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	var movie model.Netflix
	err := collection.FindOne(context.Background(), filter).Decode(&movie)
	checkNilErr(err)
	return movie
}

func GetAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movie := getOneMovies(params["id"])
	json.NewEncoder(w).Encode(movie)
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
