package main

import (
	"flag"
	"fmt"
	"os"

	"todos"
)

func addTaskCLI(task string, tasks *todos.Tasks) {
	tasks.AddTask(task)
	fmt.Println("Added task:", task)
}

func removeTaskCLI(taskID string, tasks *todos.Tasks) {
	t, err := tasks.RemoveTask(taskID)
	if err != nil {
		fmt.Println("Error removing task:", err)
		return
	}
	fmt.Println("Removed task:", t)
}

func toggleTaskCLI(taskID string, tasks *todos.Tasks) {
	t, err := tasks.GetTask(taskID)
	if err != nil {
		fmt.Println("Error finding task:", err)
		return
	}
	t.ToggleCompleted()
	fmt.Println("Toggled task:", t)
}

func editTaskCLI(taskID, newName string, tasks *todos.Tasks) {
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
	// Define flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTask := addCmd.String("task", "", "Task to add")

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeTaskID := removeCmd.String("id", "", "Task ID to remove")

	toggleCmd := flag.NewFlagSet("toggle", flag.ExitOnError)
	toggleTaskID := toggleCmd.String("id", "", "Task ID to toggle")

	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	editTaskID := editCmd.String("id", "", "Task ID to edit")
	editNewName := editCmd.String("name", "", "New name for the task")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	tasksCmd := flag.NewFlagSet("tasks", flag.ExitOnError)

	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Expected a command")
		return
	}

	cmd := os.Args[1]
	tasks := todos.Tasks{}
	tasks.Init()

	switch cmd {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addTask == "" {
			fmt.Println("Please provide a task to add using -task flag.")
			return
		}
		addTaskCLI(*addTask, &tasks)
	case "remove":
		removeCmd.Parse(os.Args[2:])
		if *removeTaskID == "" {
			fmt.Println("Please provide a task ID to remove using -id flag.")
			return
		}
		removeTaskCLI(*removeTaskID, &tasks)
	case "toggle":
		toggleCmd.Parse(os.Args[2:])
		if *toggleTaskID == "" {
			fmt.Println("Please provide a task ID to toggle using -id flag.")
			return
		}
		toggleTaskCLI(*toggleTaskID, &tasks)
	case "edit":
		editCmd.Parse(os.Args[2:])
		if *editTaskID == "" || *editNewName == "" {
			fmt.Println("Please provide a task ID and new name using -id and -name flags.")
			return
		}
		editTaskCLI(*editTaskID, *editNewName, &tasks)
	case "list":
		listCmd.Parse(os.Args[2:])
		listTasksCLI(&tasks)
	case "tasks":
		tasksCmd.Parse(os.Args[2:])
		listTasksCLI(&tasks)
	case "serve":
		serveCmd.Parse(os.Args[2:])
		todos.Serve()
	default:
		fmt.Println("Unknown command")
	}
}