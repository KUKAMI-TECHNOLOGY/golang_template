package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"main.go/database"
	"main.go/models"
)

//create user

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(models.User)

	// Parse request body into user struct
	if err := c.BodyParser(user); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err.Error()})
		return err
	}

	// Create user in the database
	if err := db.Create(&user).Error; err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err.Error()})
		return err
	}

	// Return success response
	c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has been created", "data": user})
	return nil
}

//get all user

func GetAllUser(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []models.User // Use the correct model type

	// Fetch all users and check for errors
	if err := db.Find(&users).Error; err != nil {
		// Log the error to the console
		fmt.Println("Error retrieving users:", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not retrieve users", "data": err.Error()})
	}

	// Check if users slice is empty
	if len(users) == 0 {
		fmt.Println("No users found.")
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}

	// Log the users found
	fmt.Println("Users found:", users)
	// Return the found users
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users found", "data": users})
}

// single data
func GetSingleUser(c *fiber.Ctx) error {

	db := database.DB.Db

	id := c.Params("id")
	var user models.User
	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

// delete user in db by ID
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var user models.User
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	err := db.Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}

// update a user in db
func UpdateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	var user models.User

	// Get ID from the URL parameters
	id := c.Params("id")

	// Find the user in the database by ID
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	// Define a struct that matches the fields you want to update
	var updateUserData models.User // This should match the User model

	// Parse the request body into the updateUserData struct
	if err := c.BodyParser(&updateUserData); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err.Error()})
	}

	// Update only the fields that are provided in the request body
	if updateUserData.Username != "" {
		user.Username = updateUserData.Username
	}
	if updateUserData.Email != "" {
		user.Email = updateUserData.Email
	}
	if updateUserData.Password != "" {
		user.Password = updateUserData.Password
	}

	// Save the changes to the database
	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update user", "data": err.Error()})
	}

	// Return the updated user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User updated successfully", "data": user})
}
