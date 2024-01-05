package main

import (
	"auth/connection"
	"auth/controllers"
	"auth/db"
	"auth/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// Connect to Database
	dbCon, err := connection.Connect(
		os.Getenv("dbEngine"),
		os.Getenv("dbUser"),
		os.Getenv("dbPassword"),
		os.Getenv("dbName"),
		os.Getenv("dbHost"),
		os.Getenv("dbPort"),
		os.Getenv("dbSSLMode"))

	// dbCon, err := connection.Connect(dbEngine, dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode)
	if err != nil {
		log.Fatal(err)
	}
	defer dbCon.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer dbCon.Close()

	queries := db.New(dbCon)

	userController := controllers.NewUsersController(dbCon, queries)

	app := fiber.New()

	v1 := app.Group("/")

	routes.UserRoutes(v1, userController)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong 👋!")
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://*, https://*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Accept, Content-Type, X-Requested-With, X-XSRF-TOKEN, Cookie, token",
		Next:             nil,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))
	log.Println("Server running at port 80")
	log.Fatal(app.Listen(":80"))
}
