package main

import (
	"fmt"
	"strconv"

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

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000
	if err != nil {
		fmt.Print(err)
	}
}

func initialiseDB() {
	todoDB = new(memorydb.TodoDB)
	todoDB.Initialise()
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			fmt.Printf("GET: Invalid URL /todo/%s", vars["id"])
			w.WriteHeader(404)
			fmt.Fprintf(w, "404 GET: todo ID: %s\n", id)
		}
		if todo, ok := todoDB.Read(int(id)); ok == true {
			json, _ := todo.JSON()
			w.Write(json)
		}
	})
	return r
}
