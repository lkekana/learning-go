package main

import (
	"fmt"
	"os"
	"strings"

	"todos"
)

func addTaskCLI(args []string, tasks *todos.Tasks) {
	if len(args) < 3 {
		fmt.Println("Please provide a task name.")
		return
	}
	task := strings.Join(args, " ")
	tasks.AddTask(task)
	fmt.Println("Added task:", task)
}

func removeTaskCLI(args []string, tasks *todos.Tasks) {
	if len(args) < 3 {
		fmt.Println("Please provide a task ID to remove.")
		return
	}
	taskID := args[2]
	t, err := tasks.RemoveTask(taskID)
	if err != nil {
		fmt.Println("Error removing task:", err)
		return
	}
	fmt.Println("Removed task:", t)
}

func toggleTaskCLI(args []string, tasks *todos.Tasks) {
	if len(args) < 3 {
		fmt.Println("Please provide a task ID to toggle.")
		return
	}
	taskID := args[2]
	t, err := tasks.GetTask(taskID)
	if err != nil {
		fmt.Println("Error finding task:", err)
		return
	}
	t.ToggleCompleted()
	fmt.Println("Toggled task:", t)
}

func editTaskCLI(args []string, tasks *todos.Tasks) {
	if len(args) < 4 {
		fmt.Println("Please provide a task ID and new name to edit.")
		return
	}
	taskID := args[2]
	newName := strings.Join(args[3:], " ")
	t, err := tasks.EditTask(taskID, newName)
	if err != nil {
		fmt.Println("Error editing task:", err)
		return
	}
	fmt.Println("Edited task:", t)
}

func listTasksCLI(tasks *todos.Tasks) {
	if len(tasks.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("\nTasks:")
	fmt.Println(tasks.String())
}

func main() {
    args := os.Args

	if len(args) <= 1 {
		// early return
		return
	}

	// Print all arguments
	// fmt.Printf("Args: %v (type %T)", args, args)

	cmd := args[1]
	// fmt.Println("Command:", cmd)

	tasks := todos.Tasks{}
	tasks.Init()

	switch cmd {
		case "add":
			addTaskCLI(args, &tasks)
		case "remove":
			removeTaskCLI(args, &tasks)
		case "toggle":
			toggleTaskCLI(args, &tasks)
		case "edit":
			editTaskCLI(args, &tasks)
		case "list":
			listTasksCLI(&tasks)
		case "tasks":
			listTasksCLI(&tasks)
		case "serve":
			todos.Serve()
		default:
			fmt.Println("Unknown command")
	}

    // // Print arguments excluding the program name (os.Args[1:])
    // if len(args) > 1 {
    //     fmt.Println("Arguments excluding program name:", args[1:])
    // } else {
    //     fmt.Println("No arguments provided.")
    // }
}