package cmd

import (
	models "task-manager-cli/models"
	"time"

	"github.com/google/uuid"
)

func AddTask(description string) (string, error) {
	var existingTasks = loadTasks()

	newTask := models.Task{
		ID:          uuid.New(),
		Description: description,
		Status:      models.Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	existingTasks = append(existingTasks, newTask)
	saveTasks(existingTasks)
	return newTask.ID.String(), nil
}
