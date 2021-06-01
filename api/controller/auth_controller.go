package controller

import (
	"customer/api/helper"
	"customer/domain"
	"customer/domain/model"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Service domain.AuthService
}

func (cc AuthController) Login(c *fiber.Ctx) error {
	payload := new(model.Customer)
	if err := c.BodyParser(payload); err != nil {

		return c.Status(500).JSON(helper.GenerateResponse(500, "Error when parsing the request", err.Error()))
	}

	if err := cc.Service.Login(payload.Email, payload.Password); err != nil {
		if err.Error() == "unauthorized" {
			return c.Status(400).JSON(helper.GenerateResponse(400, "Wrong username or password", nil))
		} else if err.Error() == "wrong_password" {
			return c.Status(400).JSON(helper.GenerateResponse(400, "Wrong password", nil))
		}

		return c.Status(500).JSON(helper.GenerateResponse(500, "Error when checking the customer auth", err.Error()))
	}

	return c.Status(200).JSON(helper.GenerateResponse(200, "Authentication success", nil))
}
