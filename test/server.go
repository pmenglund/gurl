package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "alive")
	fmt.Println("got request")
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	connectionString := fmt.Sprintf(":%s", port)
	http.ListenAndServe(connectionString, nil)
}
