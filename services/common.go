package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	models "task-manager-cli/models"
)

const taskFile = "./task.json"

func loadTasks() []models.Task {
	var tasks []models.Task
	file, err := os.Open(taskFile)
	if err != nil {
		return tasks // Return empty if file does not exist
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&tasks)
	return tasks
}

func saveTasks(tasks []models.Task) {
	file, err := os.Create(taskFile)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(tasks)
}
