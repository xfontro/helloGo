package main

import (
	"fmt"
	"net/http"
)

func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", getHelloWorld)
	http.HandleFunc("/health", healthCheck)
	http.ListenAndServe(":8080", nil)
}
