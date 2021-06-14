package main

import (
	"fmt"
	"golang-fifa-world-cup-web-service/data"
	"golang-fifa-world-cup-web-service/handlers"
	"net/http"
)

const (
	port=8000
)

func main() {
	data.PrintUsage()

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/winners", handlers.WinnersHandler)
	fmt.Printf("Server running on port %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
