package cmd

import models "task-manager-cli/models"

func ListTask() ([]models.Task, error) {
	var existingTasks = loadTasks()
	return existingTasks, nil
}
