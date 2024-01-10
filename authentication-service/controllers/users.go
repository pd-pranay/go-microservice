package controllers

import (
	"auth/db"
	"auth/utils"
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// UsersController is a controller and is defined here.
type UsersController struct {
	DB         *sql.DB
	Queries    *db.Queries
	HttpClient *utils.HTTPClient
}

// NewUsersController returns pointer to UsersController.
func NewUsersController(db *sql.DB, queries *db.Queries, httpClient *utils.HTTPClient) *UsersController {
	return &UsersController{
		DB:         db,
		Queries:    queries,
		HttpClient: httpClient,
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

	// log authentication
	err = uc.logRequest("authentication", fmt.Sprintf("%s logged in", user.Email))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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

func (uc *UsersController) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceURL := "http://logger-service/log"

	response, err := uc.sendPOSTRequest(jsonData, logServiceURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return errors.New("Status Code is NoN 202")
	}

	return nil
}

func (uc *UsersController) sendPOSTRequest(a any, url string) (*http.Response, error) {

	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := uc.HttpClient.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
