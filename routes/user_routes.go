package routes

import (
	"techtest/controllers"
	"techtest/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/users")

	user.Post("/register", middleware.RegisterUser) // Register user
	user.Post("/login", middleware.LoginUser)       // Login user
	user.Post("/logout", middleware.LogoutUser)     // Logout user

	user.Get("/get", middleware.Authenticate, controllers.GetUsers)      // Get all users (butuh login)
	user.Get("/:id", middleware.Authenticate, controllers.GetUser)       // Get user by ID (butuh login)
	user.Put("/:id", middleware.Authenticate, controllers.UpdateUser)    // Update user (butuh login)
	user.Delete("/:id", middleware.Authenticate, controllers.DeleteUser) // Delete user (butuh login)
	user.Get("/me", middleware.Authenticate, middleware.GetMe)           // Get user data (butuh login)
}
