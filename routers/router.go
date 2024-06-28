package routers

import (
	"github.com/gorilla/mux"
	"github.com/satti999/PractiseProject/handlers"
	"github.com/satti999/PractiseProject/middleware"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	router.HandleFunc("/items", handlers.GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", handlers.GetItem).Methods("GET")
	router.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")
	return router
}
