package main

import (
	"fmt"
	"os"
	"strconv"

	"taskmaster/todo"
)

const taskFile = "tasks.json"

func handleList(tasks []todo.Task) {
	if tasks == nil || len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		status := " "
		if task.IsDone {
			status = "x"
		}
		fmt.Printf("[%s] %d. %s\n", status, task.ID, task.Description)
	}
}

// Function for add command.
func handleAdd(tasks []todo.Task, args []string) error {
	if len(args) == 0 { return fmt.Errorf("description required") }

	// Create a new task with the provided description and a unique ID.
	newTask := todo.Task{
		ID:          len(tasks) + 1,
		Description: args,
		IsDone:      false,
	}
	// Append the new task to the existing list and save it back to the file.
	return todo.SaveTasks(taskFile, append(tasks, newTask))
}

func handleDone(tasks []todo.Task, args []string) error {
	if len(args) == 0 { return fmt.Errorf("ID required") }

	id, err := strconv.Atoi(args)

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].IsDone = true
			return todo.SaveTasks(taskFile, tasks)
		}
	}
	return fmt.Errorf("task %d not found", id)
}

func handleDelete(tasks []todo.Task, args []string) error {
	if len(args) == 0 { return fmt.Errorf("ID required") }

	id, err := strconv.Atoi(args)
	if err != nil { return fmt.Errorf("invalid ID: %v", err) }

	for i := range tasks {
		if tasks[i].ID == id {
			delTask := append(tasks[:i], tasks[i+1:]...)
			return todo.SaveTasks(taskFile, delTask)
		}
	}
	return fmt.Errorf("task %d not found", id)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: taskmaster [add|list|done|delete]")
		os.Exit(1)
	}

	// Select command from CLI arguments
	command := os.Args[1]

	// Load existing tasks from the file
	tasks, _ := todo.LoadTasks(taskFile)

	// Route the command to the right logic
	var err error
	switch command {
	case "list":
		handleList(tasks)
	case "add":
		err = handleAdd(tasks, os.Args[2:])
	case "done":
		err = handleDone(tasks, os.Args[2:])
	case "delete":
		err = handleDelete(tasks, os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
