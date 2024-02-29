package main

import (
	// "log"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func initRouter(r *mux.Router) {
	r.HandleFunc("/list", getTodos).Methods("GET")
	r.HandleFunc("/add", addTodo).Methods("POST")
	r.HandleFunc("/update/{id}", updateTodo).Methods("POST")
	r.HandleFunc("/delete/{id}", deleteTodo).Methods("DELETE")

	// http.ListenAndServe(":8000", r)
}

func main() {
	initMigratioin()

	r := mux.NewRouter()

	initRouter(r)
	http.Handle("/", r)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE"},
	})
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe("localhost:8000", handler))

	// Your CRUD operations go here
}
