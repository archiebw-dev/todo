package main

import (
	"fmt"

	"net/http"
	"os"

	"todo/internal/memorydb"

	"github.com/gorilla/mux"
)

var todoDB *memorydb.TodoDB

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	initialiseDB()
	router := setupRouter()

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}

func initialiseDB() {
	todoDB = new(memorydb.TodoDB)
	todoDB.Create()
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["id"]

		fmt.Fprintf(w, "You've requested the todo ID: %s\n", title)
	})
	return r
}
