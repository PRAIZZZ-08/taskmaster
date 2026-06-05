package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

// Task Struct
type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	IsDone bool `json:"is_done"`
}

// Function returns string representation of Task struct.
func SaveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal tasks: %v", err)
	}
	return os.WriteFile(filename, data, 0644) 
}

// Function to load tasks from a JSON file and return a slice of Task structs.
func LoadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	var tasks []Task // Declare container
	err = json.Unmarshal(data, &tasks) 
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal tasks: %v", err)
	}

	return tasks, nil 
}
