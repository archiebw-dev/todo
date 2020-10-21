package main

import (
	"fmt"

	"net/http"
	"os"

	u "todo/internal/utils"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	r := mux.NewRouter()

	r.HandleFunc("/todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["id"]

		fmt.Fprintf(w, "You've requested the todo ID: %s\n", title)
	})

	fmt.Println(port)
	u.PrintHello()

	err := http.ListenAndServe(":"+port, r) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
