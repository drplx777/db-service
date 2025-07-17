package handler

import (
	"context"
	"db-service/internal/model"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterTaskRoutes(app *fiber.App, pool *pgxpool.Pool) {
	app.Post("/create", createTask(pool))
	app.Get("/list", ListTask(pool))
	app.Get("/task/by_id", getTaskByID(pool))
	app.Delete("/delete", deleteTask(pool))
	app.Put("/update", doneTask(pool))
}

func createTask(pool *pgxpool.Pool) fiber.Handler {
	return func(c fiber.Ctx) error { // параметр теперь указатель, чтобы BodyParser доступен
		var input struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}
		if err := c.Bind().JSON(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		var task model.Task
		err := pool.QueryRow(context.Background(),
			`INSERT INTO tasks (title, description) VALUES ($1, $2)
             RETURNING id, title, description, created_at, updated_at, done_at, completed`,
			input.Title, input.Description,
		).Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt, &task.DoneAt, &task.Completed)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(task)
	}
}

func ListTask(pool *pgxpool.Pool) fiber.Handler {
	return func(c fiber.Ctx) error {
		rows, err := pool.Query(context.Background(), "SELECT id, title, description, created_at, updated_at, done_at, completed FROM tasks")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer rows.Close()

		var tasks []model.Task
		for rows.Next() {
			var t model.Task
			if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt, &t.DoneAt, &t.Completed); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			tasks = append(tasks, t)
		}
		return c.JSON(tasks)
	}
}

func deleteTask(pool *pgxpool.Pool) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Query("id")
		cmdTag, err := pool.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		if cmdTag.RowsAffected() == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
		}
		return c.JSON(fiber.Map{"status": "task deleted"})
	}
}
func doneTask(pool *pgxpool.Pool) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Query("id")
		_, err := pool.Exec(context.Background(), "UPDATE tasks SET completed = TRUE, done_at = NOW() WHERE id = $1", id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		var t model.Task
		if err := pool.QueryRow(context.Background(),
			"SELECT id, title, description, created_at, updated_at, done_at, completed FROM tasks WHERE id = $1", id,
		).Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt, &t.DoneAt, &t.Completed); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(t)
	}
}

func getTaskByID(pool *pgxpool.Pool) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Query("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing task ID"})
		}
		var t model.Task
		err := pool.QueryRow(context.Background(),
			"SELECT id, title, description, created_at, updated_at, done_at, completed FROM tasks WHERE id = $1", id,
		).Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt, &t.DoneAt, &t.Completed)
		if err != nil {
			if err.Error() == "no rows in result set" {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(t)
	}
}

func updateTask(pool *pgxpool.Pool) fiber.Handler {
	const updateTaskSQL = `
        UPDATE tasks
        SET
            title       = COALESCE(NULLIF($2, ''), title),
            description = COALESCE(NULLIF($3, ''), description),
            updated_at  = NOW()
        WHERE id = $1
        RETURNING id, title, description, created_at, updated_at, done_at, completed;
    `
	return func(c fiber.Ctx) error {
		var input struct {
			Title       string `json:"title",omitempty`
			Description string `json:"description",omitempty`
		}
		if err := c.Bind().JSON(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

		}
		id := c.Query("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing task ID"})
		}
		var t model.Task
		err := pool.QueryRow(context.Background(),
			updateTaskSQL, id, input.Title, input.Description,
		).Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt, &t.DoneAt, &t.Completed)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(t)
	}
}

// TODO: Add error handling for each function
// TODO: Add validation for input data
// TODO: add logging for each function
// TODO: add update discription and title func
