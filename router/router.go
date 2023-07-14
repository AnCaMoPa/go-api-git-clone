package router

import (
	"api-git-clone/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Check if the server is up and running
	app.Get("/health", handlers.HandleHealthCheck)

	gitclone := app.Group("/gitclone")

	//Post endpoint
	gitclone.Post("/", handlers.HandleGitAllClone)
}
