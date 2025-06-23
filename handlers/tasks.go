package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"task-manager/models"
	"task-manager/storage"
	"time"

	"github.com/google/uuid"
)

// Создание задачи
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		InputData string `json:"input_data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task := &models.Task{
		ID:        uuid.New().String(),
		Status:    models.StatusPending,
		InputData: input.InputData,
		CreatedAt: time.Now(),
	}

	storage.CreateTask(task)

	go processTask(task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": task.ID})
}

// Получение задачи
func GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task, exists := storage.GetTask(id)
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	if task.Status == models.StatusProcessing {
		duration := time.Since(task.StartedAt)
		task.Duration = duration.String()
	} else if task.Status == models.StatusCompleted || task.Status == models.StatusFailed {
		duration := (task.CompletedAt).Sub(task.StartedAt)
		task.Duration = duration.String()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Удаление задачи
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	storage.DeleteTask(id)
	w.WriteHeader(http.StatusNoContent)
}

// Оработка задачи
func processTask(task *models.Task) {
	now := time.Now()
	task.StartedAt = now
	task.Status = models.StatusProcessing

	time.Sleep(3*time.Minute + time.Duration(rand.Intn(120))*time.Second)

	now = time.Now()
	task.CompletedAt = now
	task.Status = models.StatusCompleted
	task.Result = "Processing completed for " + task.InputData
}
