package repository

import (
	"customer/domain"
	"customer/domain/model"
	"golang.org/x/net/context"
	"log"
)

// CustomerRepositoryImpl is the repository structure
type CustomerRepositoryImpl struct {
	Ctx context.Context
	DB  domain.Database
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

func (cr CustomerRepositoryImpl) Show(query ...map[string]interface{}) (*model.Customer, error) {
	var customer []model.Customer

	if err := cr.DB.Select(&customer, query...); err != nil {
		log.Println("Error when selecting", err.Error())

		return nil, err
	} else if len(customer) == 0 {
		return &model.Customer{}, nil
	}

	return &customer[0], nil
}

func (cr CustomerRepositoryImpl) Delete() bool {
	return false
}
