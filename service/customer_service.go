package service

import (
	"customer/domain"
	"customer/domain/model"
)

type CustomerServiceImpl struct {
	Repository domain.CustomerRepository
}

func (cs CustomerServiceImpl) FetchCustomer() *[]model.Customer {
	return cs.Repository.Get()
}

func (cs CustomerServiceImpl) SaveCustomer(customer *model.Customer) error {
	return cs.Repository.Insert(customer)
}
