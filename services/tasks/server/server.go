package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gdean123/full-stack-microservices/services/tasks/server/handlers"
)

func main() {
	staticFileHandler := handlers.NewStaticFileHandler(os.Args[1])
	http.Handle("/", staticFileHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
