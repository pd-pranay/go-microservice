package controllers

import (
	"auth/db"
	"auth/utils"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// UsersController is a controller and is defined here.
type UsersController struct {
	DB      *sql.DB
	Queries *db.Queries
}

// NewUsersController returns pointer to UsersController.
func NewUsersController(db *sql.DB, queries *db.Queries) *UsersController {
	return &UsersController{
		DB:      db,
		Queries: queries,
	}
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (uc *UsersController) Authenticate(c *fiber.Ctx) error {
	utils.AddHeader(c)

	body := Login{}
	// err := c.BodyParser(body)
	err := json.Unmarshal(c.BodyRaw(), &body)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	user, err := uc.Queries.GetByEmailUsers(c.Context(), body.Email)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	valid, err := uc.PasswordMatches(user.Password, body.Password)
	if err != nil || !valid {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	return c.Status(fiber.StatusAccepted).
		JSON(fiber.Map{
			"error":   false,
			"message": fmt.Sprintf("Logged in user %s", user.Email),
			"data":    db.GetByEmailUsersRow(user),
		})
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (uc *UsersController) PasswordMatches(encoded, plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(encoded), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
