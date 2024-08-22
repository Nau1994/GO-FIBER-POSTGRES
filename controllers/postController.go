package controllers

import (
	"GoFiber/config"
	"GoFiber/models"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Check if the user exists
	var user models.User
	if err := config.DB.First(&user, post.UserID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User does not exist"})
	}

	// Create the post
	if err := config.DB.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create post"})
	}

	return c.JSON(post)
}

func GetAllPosts(c *fiber.Ctx) error {
	var posts []models.Post
	config.DB.Preload("User").Find(&posts)
	return c.JSON(posts)
}
