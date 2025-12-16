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
	ID         string `json:"id"`
	Name       string `json:"name"`
	Completed  bool   `json:"completed"`
	InsertTime int64  `json:"insert_time"`
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
	Tasks []Task `json:"tasks"`
}

func (t Tasks) String() string {
	result := ""
	for _, task := range t.Tasks {
		result += fmt.Sprintf("%v\n", task.String())
	}
	return result
}

func (t *Tasks) Init() {
	tasks, err := readJSONFile()
	if err != nil {
		// fmt.Println("Error reading tasks from file:", err)
		t.Tasks = []Task{}
	} else {
		t.Tasks = tasks
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
	t.Tasks = append(t.Tasks, task)
	writeJSONFile(t.Tasks)
}

func (t *Tasks) GetTask(id string) (Task, error) {
	for i, task := range t.Tasks {
		if task.ID == id {
			return t.Tasks[i], nil
		}
	}
	return Task{}, errors.New("Task not found")
}

func (t *Tasks) GetTasks() []Task {
	return t.Tasks
}

func (t *Tasks) RemoveTask(id string) (Task, error) {
	var task Task = Task{};
	for i, task := range t.Tasks {
		if task.ID == id {
			task = t.Tasks[i]
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			writeJSONFile(t.Tasks)
			return task, nil
		}
	}
	return task, errors.New("Task not found")
}

func (t *Tasks) EditTask(id string, name string) (Task, error) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i].Name = name
			writeJSONFile(t.Tasks)
			return t.Tasks[i], nil
		}
	}
	return Task{}, errors.New("Task not found")
}

func (t *Task) ToggleCompleted() {
	t.Completed = !t.Completed
}

