package cmd

import (
	"errors"
	models "task-manager-cli/models"
	"time"
)

func MarkTask(taskId string, status models.Status) (string, error) {
	var existingTasks = loadTasks()
	found := false
	for i := range existingTasks {
		if existingTasks[i].ID.String() == taskId {
			existingTasks[i].Status = status
			existingTasks[i].UpdatedAt = time.Now().UTC() // optional
			found = true
			break
		}
	}

	if !found {
		return "", errors.New("task not found")
	}
	saveTasks(existingTasks)
	return taskId, nil
}
