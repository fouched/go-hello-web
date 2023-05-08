package main

import (
	"fmt"
	"github.com/fouched/go-course/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// pass a function to / that will handle requests
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on %s", portNumber))
	// ignore errors
	_ = http.ListenAndServe(portNumber, nil)
}
