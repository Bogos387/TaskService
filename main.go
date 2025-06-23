package main

/*
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Worlddfeff!")
	})
	http.ListenAndServe(":8080", nil)
}
*/

import (
	"fmt"
	"log"
	"net/http"
	"task-manager/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World 2!")
	})
	/*
		// Регистрируем обработчики
		http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				handlers.CreateTask(w, r)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		})

		http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				handlers.GetTask(w, r)
			case http.MethodDelete:
				handlers.DeleteTask(w, r)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		})*/

	http.HandleFunc("/tasks/create", handlers.CreateTask)
	http.HandleFunc("/tasks/get", handlers.GetTask)
	http.HandleFunc("/tasks/delete", handlers.DeleteTask)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
