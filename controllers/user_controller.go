package controllers

import (
	"techtest/configs"
	"techtest/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Get All Users
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	configs.DB.Find(&users)
	return c.JSON(users)
}

// Get User By ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := configs.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// Create User
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	user.Password = string(hashedPassword)

	if err := configs.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}
	return c.JSON(user)
}

// Update User
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if err := configs.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Hash Password if provided
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		user.Password = string(hashedPassword)
	}

	configs.DB.Save(&user)
	return c.JSON(user)
}

// Delete User
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := configs.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	configs.DB.Delete(&user)
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
