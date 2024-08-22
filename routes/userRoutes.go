package routes

import (
	"GoFiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userGroup := app.Group("/users")

	userGroup.Post("/", controllers.CreateUser)
	userGroup.Get("/", controllers.GetAllUsers)
	userGroup.Post("/transfer", controllers.TransferFunds)
	userGroup.Get("/with-posts", controllers.GetUsersWithPosts)
	userGroup.Get("/post-counts", controllers.GetUserPostCounts)
	userGroup.Get("/recent-posts", controllers.GetUsersWithRecentPosts)
	userGroup.Get("/:id", controllers.GetUser)
}
