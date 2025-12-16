package todos

import (
	"github.com/gofiber/fiber/v2"
)

var taskManager = Tasks{}

func Serve() {
	taskManager.Init()
	app := fiber.New()

	app.Get("/tasks", getTasks)
	app.Get("/tasks/:id", getTask)
	app.Post("/tasks", addTask)
	app.Delete("/tasks/:id", removeTask)
	app.Patch("/tasks/:id/toggle", toggleTask)

	app.Listen("localhost:8080")
}

func getTasks(c *fiber.Ctx) error {
	return c.JSON(taskManager.Tasks)
}

func getTask(c *fiber.Ctx) error {
	taskID := c.Params("id")
	t, err := taskManager.GetTask(taskID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}
	return c.JSON(t)
}

func addTask(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`;
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	taskManager.AddTask(body.Name)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Task added",
	})
}

func removeTask(c *fiber.Ctx) error {
	taskID := c.Params("id")
	t, err := taskManager.RemoveTask(taskID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Removed task",
		"task":    t,
	})
}

func toggleTask(c *fiber.Ctx) error {
	taskID := c.Params("id")
	t, err := taskManager.GetTask(taskID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}
	t.ToggleCompleted()
	return c.JSON(fiber.Map{
		"message": "Toggled task",
		"task":    t,
	})
}

