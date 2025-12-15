package todos

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/xeonx/timeago"
)

type Task struct {
	ID string;
	Name string;
	Completed bool;
	InsertTime int64;
}

func (t Task) String() string {
	var checkmark string;
	if t.Completed {
		checkmark = "☑︎"
	} else {
		checkmark = "☐"
	}
	return fmt.Sprintf("%v %v\t(%v\t| %v)", checkmark, t.Name, timeago.English.Format(time.Unix(t.InsertTime, 0)), t.ID)
}

type Tasks struct {
	tasks []Task;
}

func (t Tasks) String() string {
	result := ""
	for _, task := range t.tasks {
		result += fmt.Sprintf("%v\n", task.String())
	}
	return result
}

func (t *Tasks) Init() {
	tasks, err := readJSONFile()
	if err != nil {
		fmt.Println("Error reading tasks from file:", err)
		t.tasks = []Task{}
	} else {
		t.tasks = tasks
	}
}

func readJSONFile() ([]Task, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil, err
	}
	if len(data) > 0 {
		var tasks []Task
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			return nil, err
		}
		return tasks, nil
	}
	return []Task{}, nil
}

func writeJSONFile(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tasks) AddTask(name string) {
	uuid := uuid.New()
	id := uuid.String()
	task := Task{
		ID: id,
		Name: name,
		Completed: false,
		InsertTime: time.Now().Unix(),
	}
	t.tasks = append(t.tasks, task)
	writeJSONFile(t.tasks)
}

func (t *Tasks) GetTask(id string) (Task, error) {
	for i, task := range t.tasks {
		if task.ID == id {
			return t.tasks[i], nil
		}
	}
	return Task{}, errors.New("Task not found")
}

func (t *Tasks) GetTasks() []Task {
	return t.tasks
}

func (t *Tasks) RemoveTask(id string) (Task, error) {
	var task Task = Task{};
	for i, task := range t.tasks {
		if task.ID == id {
			task = t.tasks[i]
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			writeJSONFile(t.tasks)
			return task, nil
		}
	}
	return task, errors.New("Task not found")
}

func (t *Tasks) EditTask(id string, name string) (Task, error) {
	for i, task := range t.tasks {
		if task.ID == id {
			t.tasks[i].Name = name
			writeJSONFile(t.tasks)
			return t.tasks[i], nil
		}
	}
	return Task{}, errors.New("Task not found")
}

func (t *Task) ToggleCompleted() {
	t.Completed = !t.Completed
}

