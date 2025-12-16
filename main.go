package main

import (
	"fmt"
	"os"
	"strings"

	"todos"
)

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
			if len(args) < 3 {
				fmt.Println("Please provide a task name.")
				return
			}
			task := strings.Join(args[2:], " ")
			tasks.AddTask(task)
			fmt.Println("Added task:", task)
		case "remove":
			if len(args) < 3 {
				fmt.Println("Please provide a task ID to remove.")
			}
			taskID := args[2]
			t, err := tasks.RemoveTask(taskID)
			if err != nil {
				fmt.Println("Error removing task:", err)
				return
			}
			fmt.Println("Removed task:", t)
		case "toggle":
			if len(args) < 3 {
				fmt.Println("Please provide a task ID to toggle.")
			}
			taskID := args[2]
			t, err := tasks.GetTask(taskID)
			if err != nil {
				fmt.Println("Error finding task:", err)
				return
			}
			t.ToggleCompleted()
			fmt.Println("Toggled task:", t)
		case "edit":
			if len(args) < 3 {
				fmt.Println("Please provide a task ID to edit.")
			}
			taskID := args[2]
			t, err := tasks.GetTask(taskID)
			if err != nil {
				fmt.Println("Error finding task:", err)
				return
			}
			newName := strings.Join(args[3:], " ")
			t, err = tasks.EditTask(t.ID, newName)
			if err != nil {
				fmt.Println("Error editing task:", err)
				return
			}
			fmt.Println("Edited task name to:", newName)
		case "list":
			if len(tasks.Tasks) == 0 {
				fmt.Println("No tasks found.")
				return
			}
			fmt.Println("\nTasks:")
			fmt.Println(tasks.String())
		case "tasks":
			if len(tasks.Tasks) == 0 {
				fmt.Println("No tasks found.")
				return
			}
			fmt.Println("\nTasks:")
			fmt.Println(tasks.String())
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