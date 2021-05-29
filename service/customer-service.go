package service

import (
	"customer/domain"
	"customer/domain/model"
)

type CustomerService struct {
	Repository domain.ICustomerRepository
}

func (cs CustomerService) AuthenticateUser() bool {
	return false
}

func (cs CustomerService) FetchCustomer() *[]model.Customer {
	return cs.Repository.Get()
}

func (cs CustomerService) SaveCustomer(customer *model.Customer) error {
	return cs.Repository.Insert(customer)
}
