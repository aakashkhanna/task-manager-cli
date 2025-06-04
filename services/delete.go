package cmd

import (
	"errors"
)

func DeleteTask(taskId string) (string, error) {
	var existingTasks = loadTasks()
	found := false
	taskIndex := 0
	for idx, task := range existingTasks {
		if task.ID.String() == taskId {
			found = true
			taskIndex = idx
		}
	}

	if !found {
		return "", errors.New("task not found")
	}
	existingTasks = append(existingTasks[:taskIndex], existingTasks[taskIndex+1:]...)
	saveTasks(existingTasks)
	return taskId, nil
}
