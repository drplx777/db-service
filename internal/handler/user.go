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
		var user model.User
		if err := c.Bind().JSON(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Хеширование пароля
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
		}
		user.Password = string(hashedPassword)

		err = pool.QueryRow(context.Background(),
			`INSERT INTO users (name, sumame, middlename, login, roleID, password) 
			 VALUES ($1, $2, $3, $4, $5, $6) 
			 RETURNING id`,
			user.Name, user.Sumame, user.Middlename, user.Login, user.RoleID, user.Password,
		).Scan(&user.ID)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Не возвращаем пароль в ответе
		user.Password = ""
		return c.Status(fiber.StatusCreated).JSON(user)
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
		).Scan(&user.ID, &user.Name, &user.Sumame, &user.Middlename, &user.Login, &user.RoleID, &user.Password)

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}

		return c.JSON(user)
	}
}
