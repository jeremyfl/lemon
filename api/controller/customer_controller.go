package controller

import (
	"customer/domain"
	"customer/domain/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

type CustomerController struct {
	Service domain.CustomerService
}

func (cc CustomerController) Get(c *fiber.Ctx) error {
	payload := model.Customer{
		ID:        1,
		FirstName: "Sam",
		LastName:  "Smith",
		Password:  "",
		Age:       30,
	}

	if err := cc.Service.SaveCustomer(&payload); err != nil {
		log.Println("Error when inserting data to the repo", err)
	}

	customer := cc.Service.FetchCustomer()

	return c.JSON(customer)
}

func (cc CustomerController) Post(c *fiber.Ctx) error {
	return nil
	//cc.Service.SaveCustomer()
}
