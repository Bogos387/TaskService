package storage

import (
	"sync"
	"task-manager/models"
)

var (
	tasks     = make(map[string]*models.Task)
	tasksLock sync.RWMutex
)

func CreateTask(task *models.Task) {
	tasksLock.Lock()
	defer tasksLock.Unlock()
	tasks[task.ID] = task
}

func GetTask(id string) (*models.Task, bool) {
	tasksLock.RLock()
	defer tasksLock.RUnlock()
	task, exists := tasks[id]
	return task, exists
}

func DeleteTask(id string) {
	tasksLock.Lock()
	defer tasksLock.Unlock()
	delete(tasks, id)
}
