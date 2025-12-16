package todos

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestApp() *fiber.App {
	taskManager = Tasks{}
	taskManager.Init()
	app := fiber.New()

	app.Get("/tasks", getTasks)
	app.Get("/tasks/:id", getTask)
	app.Post("/tasks", addTask)
	app.Delete("/tasks/:id", removeTask)
	app.Patch("/tasks/:id", toggleTask)

	return app
}

func TestHTTPGetTasks(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHTTPAddTask(t *testing.T) {
	app := setupTestApp()

	body := map[string]string{"name": "Test Task"}
	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestHTTPGetTask(t *testing.T) {
	app := setupTestApp()

	// Add a task first
	taskManager.AddTask("Test Task")
	taskID := taskManager.Tasks[0].ID

	req := httptest.NewRequest(http.MethodGet, "/tasks/"+taskID, nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHTTPGetTaskNotFound(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodGet, "/tasks/null", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestHTTPRemoveTask(t *testing.T) {
	app := setupTestApp()

	// Add a task first
	taskManager.AddTask("Test Task")
	taskID := taskManager.Tasks[0].ID

	req := httptest.NewRequest(http.MethodDelete, "/tasks/"+taskID, nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHTTPRemoveTaskNotFound(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodDelete, "/tasks/null", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestHTTPToggleTask(t *testing.T) {
	app := setupTestApp()

	// Add a task first
	taskManager.AddTask("Test Task")
	taskID := taskManager.Tasks[0].ID

	req := httptest.NewRequest(http.MethodPatch, "/tasks/"+taskID, nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHTTPToggleTaskNotFound(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest(http.MethodPatch, "/tasks/null", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}