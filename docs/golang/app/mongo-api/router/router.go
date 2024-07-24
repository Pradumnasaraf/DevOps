package router

import (
	"github.com/gorilla/mux"
	"github.com/pradumnasaraf/mongo-api/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.ServeHomepage).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controller.GetAMovie).Methods("GET")
	router.HandleFunc("/api/movies", controller.GetMyAllMovie).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteMyAllMovie).Methods("DELETE")

	return router
}
