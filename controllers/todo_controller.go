package controllers

import (
	"techtest/configs"
	"techtest/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// Buat To-Do List
func CreateTodoList(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*models.JWTClaims) // Menggunakan struct custom JWTClaims
	userID := claims.UserID                   // Mengakses userId dari struct custom

	var todoList models.TodoList
	if err := c.BodyParser(&todoList); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	todoList.UserID = userID
	if err := configs.DB.Create(&todoList).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create To-Do List"})
	}

	return c.JSON(todoList)
}

// Tambah Todo ke dalam Todo List
func AddTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*models.JWTClaims) // Menggunakan struct custom JWTClaims
	userID := claims.UserID                   // Mengakses userId dari struct custom

	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	var todoList models.TodoList
	if err := configs.DB.First(&todoList, "id = ? AND user_id = ?", todo.TodoListID, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "To-Do List not found"})
	}

	if err := configs.DB.Create(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to add To-Do"})
	}

	return c.JSON(todo)
}

// Dapatkan semua Todo Lists milik user
func GetTodoLists(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*models.JWTClaims) // Menggunakan struct custom JWTClaims
	userID := claims.UserID                   // Mengakses userId dari struct custom

	var todoLists []models.TodoList
	if err := configs.DB.Preload("Todos").Where("user_id = ?", userID).Find(&todoLists).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to retrieve To-Do Lists"})
	}

	return c.JSON(todoLists)
}

// Detail Checklist berdasarkan ID
func GetTodoListByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var todoList models.TodoList

	if err := configs.DB.Preload("Todos").First(&todoList, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Checklist not found"})
	}

	return c.JSON(todoList)
}

// Detail Item berdasarkan ID
func GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo

	if err := configs.DB.First(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "To-Do item not found"})
	}

	return c.JSON(todo)
}

// Edit To-Do
func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	var todo models.Todo
	if err := configs.DB.First(&todo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "To-Do not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to retrieve To-Do"})
	}

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	if err := configs.DB.Save(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update To-Do"})
	}

	return c.JSON(todo)
}

// Hapus To-Do
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := configs.DB.Delete(&models.Todo{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete To-Do"})
	}

	return c.JSON(fiber.Map{"message": "To-Do deleted successfully"})
}

func DeleteTodoList(c *fiber.Ctx) error {
	id := c.Params("id")

	// Hapus semua item di dalam checklist
	if err := configs.DB.Where("todo_list_id = ?", id).Delete(&models.Todo{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete checklist items"})
	}

	// Hapus checklist itu sendiri
	if err := configs.DB.Delete(&models.TodoList{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete checklist"})
	}

	return c.JSON(fiber.Map{"message": "Checklist deleted successfully"})
}

// Tandai sebagai selesai
func CompleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	var todo models.Todo
	if err := configs.DB.First(&todo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "To-Do not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to retrieve To-Do"})
	}

	todo.Completed = true
	if err := configs.DB.Save(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update To-Do status"})
	}

	return c.JSON(todo)
}
