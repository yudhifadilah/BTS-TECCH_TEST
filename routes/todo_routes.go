package routes

import (
	"techtest/controllers"
	"techtest/middleware"

	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(app *fiber.App) {
	api := app.Group("/api/todo")

	// Proteksi dengan JWT
	api.Use(middleware.Authenticate)

	api.Post("/list", controllers.CreateTodoList)        // Buat daftar To-Do
	api.Get("/lists", controllers.GetTodoLists)          // Dapatkan semua daftar To-Do
	api.Get("/lists/:id", controllers.GetTodoListByID)   // Dapatkan Detail Checklist berdasarkan ID
	api.Get("/:id", controllers.GetTodoByID)             // Dapatkan Detail Item berdasarkan ID
	api.Post("/", controllers.AddTodo)                   // Tambah To-Do dalam daftar
	api.Put("/:id", controllers.UpdateTodo)              // Edit To-Do
	api.Delete("/:id", controllers.DeleteTodo)           // Hapus To-Do
	api.Delete("/lists/:id", controllers.DeleteTodoList) // Hapus To-Do List-check
	api.Patch("/:id/complete", controllers.CompleteTodo) // Tandai sebagai selesai
}
