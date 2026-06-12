package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.Ldate|log.Ltime|log.Lshortfile)

	srv := server.NewServer(logger)

	logger.Println("Starting server on :8080")

	err := srv.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
