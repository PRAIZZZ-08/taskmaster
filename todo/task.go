package todo

import (
	"encoding/json"
	"os"
	"fmt"
)

// Task Struct
type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	isDone bool `json:"is_done"`
}

// Function returns string representation of Task struct.
func SaveTasks(filename string, tasks []Task) error {
taskData, err := json.MarshalIndent(tasks, "", "  ")

if err != nil {
	return fmt.Errorf("Error: Unable to marshal JSON data", err)
}
return os.WriteFile(filename, taskData)
}

// Function to load tasks from a JSON file and return a slice of Task structs.
func LoadTasks(filename string) ([]Task, error) {
jsonData, err := os.ReadFile(filename)

if err != nil {
	return fmt.Errorf("Error: Unable to read file", err)
}
	return json.Unmarshal(jsonData, &tasks)
}