package controllers

import (
	"GoFiber/config"
	"GoFiber/models"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TransferFunds(c *fiber.Ctx) error {
	type TransferRequest struct {
		FromUserID uint `json:"from_user_id"`
		ToUserID   uint `json:"to_user_id"`
		Amount     int  `json:"amount"`
	}

	var req TransferRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	var fromUser, toUser models.User

	err := config.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.First(&fromUser, req.FromUserID).Error; err != nil {
			return err
		}

		if err := tx.First(&toUser, req.ToUserID).Error; err != nil {
			return err
		}

		if fromUser.Balance < req.Amount {
			return errors.New("insufficient funds")
		}

		fromUser.Balance -= req.Amount
		toUser.Balance += req.Amount

		if err := tx.Save(&fromUser).Error; err != nil {
			return err
		}

		if err := tx.Save(&toUser).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "transfer successful",
		"fromUser": fromUser, "toUser": toUser})
}

func GetUsersWithPosts(c *fiber.Ctx) error {
	var users []models.User
	config.DB.Preload("Posts").Find(&users)
	return c.JSON(users)
}

func GetUserPostCounts(c *fiber.Ctx) error {
	var result []struct {
		UserID    uint
		PostCount int
	}

	config.DB.Model(&models.Post{}).Select("user_id, count(*) as post_count").Group("user_id").Scan(&result)
	return c.JSON(result)
}

func GetUsersWithRecentPosts(c *fiber.Ctx) error {
	// Define the query using raw SQL
	query := `
        SELECT DISTINCT u.*
        FROM users u
        JOIN posts p ON u.id = p.user_id
        WHERE p.created_at > ?
    `

	// Execute the query
	var users []models.User
	err := config.DB.Raw(query, time.Now().Add(-7*24*time.Hour)).Find(&users).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not fetch users"})
	}

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	config.DB.Preload("Posts").Find(&users)
	return c.JSON(users)
}
