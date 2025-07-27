package handler

import (
	"context"
	"db-service/internal/model"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserRoutes(app *fiber.App, pool *pgxpool.Pool) {
	app.Post("/user/register", registerUser(pool))
	app.Get("/user/by-login", getUserByLogin(pool))
}

func registerUser(pool *pgxpool.Pool) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req struct {
			Name       string `json:"name"`
			Surname    string `json:"surname"`
			Middlename string `json:"middlename,omitempty"`
			Login      string `json:"login"`
			RoleID     int    `json:"roleID"`
			Password   string `json:"password"`
		}

		if err := c.Bind().JSON(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
		}

		// Хеширование пароля
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not hash password"})
		}

		var userID int
		err = pool.QueryRow(context.Background(),
			`INSERT INTO users (name, surname, middlename, login, roleID, password) 
             VALUES ($1, $2, $3, $4, $5, $6) 
             RETURNING id`,
			req.Name, req.Surname, req.Middlename, req.Login, req.RoleID, string(hashedPassword),
		).Scan(&userID)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": userID})
	}
}

func getUserByLogin(pool *pgxpool.Pool) fiber.Handler {
	return func(c fiber.Ctx) error {
		login := c.Query("login")
		if login == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing login"})
		}

		var user model.User
		err := pool.QueryRow(context.Background(),
			`SELECT id, name, sumame, middlename, login, roleID, password 
			 FROM users WHERE login = $1`, login,
		).Scan(&user.ID, &user.Name, &user.Surname, &user.Middlename, &user.Login, &user.RoleID, &user.Password)

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}

		return c.JSON(user)
	}
}
