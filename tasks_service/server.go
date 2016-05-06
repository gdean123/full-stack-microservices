package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gdean123/full-stack-microservices/tasks_service/handlers"
)

func main() {
	staticFileHandler := handlers.NewStaticFileHandler(os.Args[1])
	http.Handle("/", staticFileHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
