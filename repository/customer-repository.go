package repository

import (
	"customer/domain/model"
	"golang.org/x/net/context"
	"log"
)

// CustomerRepository is the repository structure
type CustomerRepository struct {
	Ctx context.Context
	DB  Database
}

func (cr CustomerRepository) Get() *[]model.Customer {
	var customers []model.Customer

	if err := cr.DB.Select(&customers); err != nil {
		log.Println("Error when selecting", err.Error())
	}

	return &customers
}

func (cr CustomerRepository) Insert(customer *model.Customer) error {
	_, err := cr.DB.Insert(customer)
	return err
}

func (cr CustomerRepository) Show() *model.Customer {
	return &model.Customer{
		FirstName: "",
		LastName:  "",
		Password:  "",
		Age:       0,
	}
}

func (cr CustomerRepository) Delete() bool {
	return false
}
