package main

import (
	"log"
	"net/http"
	"ws/internal/handlers"
)

func main() {

	mux := routes()
	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()
	log.Println("Starting web server on port 8080")
	_ = http.ListenAndServe("127.0.0.1:8080", mux)
}
