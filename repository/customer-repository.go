package repository

import (
	"customer/domain/model"
	"customer/internal"
	"golang.org/x/net/context"
	"log"
)

// CustomerRepositoryImpl is the repository structure
type CustomerRepositoryImpl struct {
	Ctx context.Context
	DB  internal.Database
}

func (cr CustomerRepositoryImpl) Get() *[]model.Customer {
	var customers []model.Customer

	if err := cr.DB.Select(&customers); err != nil {
		log.Println("Error when selecting", err.Error())
	}

	return &customers
}

func (cr CustomerRepositoryImpl) Insert(customer *model.Customer) error {
	_, err := cr.DB.Insert(customer)
	return err
}

func (cr CustomerRepositoryImpl) Show() *model.Customer {
	return &model.Customer{
		FirstName: "",
		LastName:  "",
		Password:  "",
		Age:       0,
	}
}

func (cr CustomerRepositoryImpl) Delete() bool {
	return false
}
