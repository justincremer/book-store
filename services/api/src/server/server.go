package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/justincremer/go-orm/src/database"
	"github.com/justincremer/go-orm/src/models/book"
	"github.com/justincremer/go-orm/src/models/user"
)

const (
	host     string = "localhost"
	dbport   string = "5432"
	username string = "admin"
	password string = "admin"
	dbname   string = "admin"
	sslmode  string = "disable"
	timezone string = "America/Los_Angeles"
)

func Create(port string, config fiber.Config) *fiber.App {
	// Initialize Server
	app := fiber.New(config)

	// Connect to db
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", host, dbport, username, password, dbname, sslmode, timezone)
	conn := database.Connect(dsn)

	// Run migrations
	conn.AutoMigrate(&book.Book{})
	conn.AutoMigrate(&user.User{})

	// Middleware

	// Use Endpoints
	app.Get("/", welcome)
	bookRouter(app)
	userRouter(app)

	return app
}
