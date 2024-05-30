package controllers

import (
	"apique/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	var queue []models.Queue

	models.DB.Find(&queue)

	return c.Status(fiber.StatusOK).JSON(queue)

}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")

	var queue models.Queue

	if err := models.DB.Model(&queue).Where("id = ?", id).First(&queue).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Kamar dengan id " + id + " tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server sedang mengalami gangguan",
		})
	}

	return c.JSON(queue)
}

func Create(c *fiber.Ctx) error {

	var queue models.Queue

	if err := c.BodyParser(&queue); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	
	queue.Created = time.Now()
	queue.Updated = time.Now()

	if err := models.DB.Create(&queue).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Antrian berhasil dibuat",
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var queue models.Queue

	if err := c.BodyParser(&queue); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	queue.Updated = time.Now()

	if models.DB.Where("id = ?", id).Updates(&queue).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Kamar dengan id " + id + " tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diperbaharui",
	})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	var queue models.Queue

	if models.DB.Where("id = ?", id).Delete(&queue).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Kamar dengan id " + id + " tidak dapat dihapus",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus data kamar",
	})
}