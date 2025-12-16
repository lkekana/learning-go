package todos

import (
	"os"
	"testing"
	"time"
)

func setup() *Tasks {
	tasks := &Tasks{}
	tasks.Init()
	return tasks
}

func teardown() {
	os.Remove("tasks.json")
}

func TestAddTask(t *testing.T) {
	tasks := setup()
	defer teardown()

	tasks.AddTask("Test Task")
	if len(tasks.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks.Tasks))
	}

	if tasks.Tasks[0].Name != "Test Task" {
		t.Errorf("Expected task name 'Test Task', got '%s'", tasks.Tasks[0].Name)
	}
}

func TestGetTask(t *testing.T) {
	tasks := setup()
	defer teardown()

	tasks.AddTask("Test Task")
	task := tasks.Tasks[0]

	retrievedTask, err := tasks.GetTask(task.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if retrievedTask.ID != task.ID {
		t.Errorf("Expected task ID '%s', got '%s'", task.ID, retrievedTask.ID)
	}
}

func TestGetTaskNotFound(t *testing.T) {
	tasks := setup()
	defer teardown()

	_, err := tasks.GetTask("null")
	if err == nil || err.Error() != "Task not found" {
		t.Errorf("Expected 'Task not found' error, got '%v'", err)
	}
}

func TestRemoveTask(t *testing.T) {
	tasks := setup()
	defer teardown()

	tasks.AddTask("Test Task")
	task := tasks.Tasks[0]

	removedTask, err := tasks.RemoveTask(task.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if removedTask.ID != task.ID {
		t.Errorf("Expected removed task ID '%s', got '%s'", task.ID, removedTask.ID)
	}

	if len(tasks.Tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks.Tasks))
	}
}

func TestRemoveTaskNotFound(t *testing.T) {
	tasks := setup()
	defer teardown()

	_, err := tasks.RemoveTask("null")
	if err == nil || err.Error() != "Task not found" {
		t.Errorf("Expected 'Task not found' error, got '%v'", err)
	}
}

func TestEditTask(t *testing.T) {
	tasks := setup()
	defer teardown()

	tasks.AddTask("Old Task Name")
	task := tasks.Tasks[0]

	editedTask, err := tasks.EditTask(task.ID, "New Task Name")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if editedTask.Name != "New Task Name" {
		t.Errorf("Expected task name 'New Task Name', got '%s'", editedTask.Name)
	}
}

func TestEditTaskNotFound(t *testing.T) {
	tasks := setup()
	defer teardown()

	_, err := tasks.EditTask("null", "New Task Name")
	if err == nil || err.Error() != "Task not found" {
		t.Errorf("Expected 'Task not found' error, got '%v'", err)
	}
}

func TestToggleCompleted(t *testing.T) {
	task := Task{
		ID:         "1",
		Name:       "Test Task",
		Completed:  false,
		InsertTime: time.Now().Unix(),
	}

	task.ToggleCompleted()
	if !task.Completed {
		t.Errorf("Expected task to be completed, got not completed")
	}

	task.ToggleCompleted()
	if task.Completed {
		t.Errorf("Expected task to be not completed, got completed")
	}
}