package service

import (
	"customer/domain/mocks"
	"customer/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var customerRepository = mocks.CustomerRepository{Mock: mock.Mock{}}
var customerService = CustomerServiceImpl{Repository: &customerRepository}

func TestCustomerServiceImpl_FetchCustomer(t *testing.T) {
	expectedCustomer := []model.Customer{
		{
			ID:        1,
			FirstName: "Jeremiah",
			LastName:  "Ferdinand",
			Password:  "123",
			Age:       26,
		},
		{
			ID:        2,
			FirstName: "Ben",
			LastName:  "Chill",
			Password:  "123",
			Age:       26,
		},
	}

	customerRepository.Mock.On("Get").Return(expectedCustomer)
	result := customerService.FetchCustomer()

	assert.NotNil(t, *result)
	assert.Equal(t, expectedCustomer, *result)
}

func TestCustomerServiceImpl_SaveCustomer(t *testing.T) {
}
