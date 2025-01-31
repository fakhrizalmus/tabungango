package controllers

import (
	"github.com/fakhrizalmus/tabungango/config"
	model "github.com/fakhrizalmus/tabungango/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var (
		req model.User
	)

	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	if err := config.DB.Create(&req).Error; err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": true,
		"data":   req,
	})
}

func Tabung(c *fiber.Ctx) error {
	type Request struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int64  `json:"nominal"`
	}

	var (
		req  Request
		user model.User
	)

	err := c.BodyParser(&req)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	if err := config.DB.Where("no_rekening = ?", req.NoRekening).First(&user).Error; err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}
	user.Saldo = user.Saldo + req.Nominal
	config.DB.Save(&user)
	return c.JSON(fiber.Map{
		"status": true,
		"data":   user.Saldo,
	})
}

func Tarik(c *fiber.Ctx) error {
	type Request struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int64  `json:"nominal"`
	}

	var (
		req  Request
		user model.User
	)

	err := c.BodyParser(&req)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	if err := config.DB.Where("no_rekening = ?", req.NoRekening).First(&user).Error; err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}
	if user.Saldo < req.Nominal {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  false,
			"message": "Saldo tidak mencukupi",
		})
	}
	user.Saldo = user.Saldo - req.Nominal
	config.DB.Save(&user)
	return c.JSON(fiber.Map{
		"status": true,
		"data":   user.Saldo,
	})
}
