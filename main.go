package main

import (
	"fmt"

	"net/http"
	"os"

	u "todo/internal/utils"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	fmt.Println(port)
	u.PrintHello()

	err := http.ListenAndServe(":"+port, nil) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
