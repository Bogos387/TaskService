package main

import (
	"log"
	"net/http"
	"task-manager/handlers"
)

func main() {
	http.HandleFunc("/tasks/create", handlers.CreateTask)
	http.HandleFunc("/tasks/get", handlers.GetTask)
	http.HandleFunc("/tasks/delete", handlers.DeleteTask)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
