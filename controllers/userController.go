package controllers

import (
	"fmt"
	"time"

	"github.com/fakhrizalmus/tabungango/config"
	model "github.com/fakhrizalmus/tabungango/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var (
		req  model.User
		user model.User
	)

	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": false,
			"remark": err.Error(),
		})
	}

	currentTime := time.Now()
	date := currentTime.Format("012006")
	var count int64
	startMounth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)
	endMounth := startMounth.AddDate(0, 1, -1)

	config.DB.Model(&user).Where("created_at BETWEEN ? AND ?", startMounth, endMounth).Count(&count)

	sequence := count + 1
	req.NoRekening = fmt.Sprintf("%s%04d", date, sequence)

	if err := config.DB.Create(&req).Error; err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status": false,
			"remark": "NIK atau No HP sudah digunakan!",
		})
	}

	return c.JSON(fiber.Map{
		"status":      true,
		"no_rekening": req.NoRekening,
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
			"status": false,
			"remark": err.Error(),
		})
	}

	if err := config.DB.Where("no_rekening = ?", req.NoRekening).First(&user).Error; err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status": false,
			"remark": "No Rekening tidak ditemukan!",
		})
	}
	user.Saldo = user.Saldo + req.Nominal
	config.DB.Save(&user)
	return c.JSON(fiber.Map{
		"status": true,
		"saldo":  user.Saldo,
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
			"status": false,
			"remark": err.Error(),
		})
	}

	if err := config.DB.Where("no_rekening = ?", req.NoRekening).First(&user).Error; err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status": false,
			"remark": "No Rekening tidak ditemukan!",
		})
	}
	if user.Saldo < req.Nominal {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status": false,
			"remark": "Saldo tidak mencukupi",
		})
	}
	user.Saldo = user.Saldo - req.Nominal
	config.DB.Save(&user)
	return c.JSON(fiber.Map{
		"status": true,
		"saldo":  user.Saldo,
	})
}

func Saldo(c *fiber.Ctx) error {
	noRekening := c.Params("no_rekening")
	var user model.User

	if err := config.DB.Where("no_rekening = ?", noRekening).First(&user).Error; err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status": false,
			"remark": "No Rekening tidak ditemukan!",
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"saldo":  user.Saldo,
	})
}
