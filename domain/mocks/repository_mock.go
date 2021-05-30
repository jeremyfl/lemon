package mocks

import (
	"customer/domain/model"
	mock "github.com/stretchr/testify/mock"
)

type CustomerRepository struct {
	mock.Mock
}

func (c *CustomerRepository) Get() *[]model.Customer {
	arguments := c.Mock.Called()

	customers := arguments.Get(0).([]model.Customer)

	return &customers
}

func (c *CustomerRepository) Insert(customer *model.Customer) error {
	panic("implement me")
}

func (c *CustomerRepository) Show() *model.Customer {
	panic("implement me")
}

func (c *CustomerRepository) Delete() bool {
	panic("implement me")
}
