package controller

import (
	"customer/api/helper"
	"customer/domain"
	"customer/domain/model"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type CustomerController struct {
	Service domain.CustomerService
}

func (cc CustomerController) Get(c *fiber.Ctx) error {
	customer := cc.Service.FetchCustomer()

	return c.JSON(helper.GenerateResponse(200, "", customer))
}

func (cc CustomerController) Post(c *fiber.Ctx) error {
	payload := new(model.Customer)
	if err := c.BodyParser(payload); err != nil {

		return c.Status(500).JSON(helper.GenerateResponse(500, "Error when parsing the request", nil))
	}

	if err := cc.Service.SaveCustomer(payload); err != nil {

		log.Println("Error when inserting data to the repo", err)

		return c.Status(500).JSON(helper.GenerateResponse(500, "Error when inserting data to the repo", nil))
	}

	return c.Status(201).JSON(helper.GenerateResponse(201, "Data created", payload))
}

func (cc CustomerController) Show(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	customer := cc.Service.ShowCustomer(id)

	if customer.ID == 0 {
		return c.Status(404).JSON(helper.GenerateResponse(404, "We can't found that customer_id", nil))
	}

	return c.Status(200).JSON(helper.GenerateResponse(200, "", customer))
}
