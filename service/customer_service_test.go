package service

import (
	"customer/domain/mocks"
	"customer/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	customerRepository = mocks.CustomerRepository{Mock: mock.Mock{}}
	customerService    = CustomerServiceImpl{Repository: &customerRepository}
)

func TestCustomerServiceImpl_FetchCustomer(t *testing.T) {
	expectedResult := &[]model.Customer{
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

	customerRepository.Mock.On("Get").Return(expectedResult)
	result := customerService.FetchCustomer()

	assert.NotNil(t, *result)
	assert.Equal(t, *expectedResult, *result)
}

func TestCustomerServiceImpl_SaveCustomer(t *testing.T) {
	payload := &model.Customer{
		ID:        1,
		FirstName: "Jeremiah",
		LastName:  "Ferdinand",
		Password:  "123",
		Age:       26,
	}

	customerRepository.Mock.On("Insert", payload).Return(nil)

	result := customerService.SaveCustomer(payload)

	assert.Nil(t, result)
}

func TestCustomerServiceImpl_ShowCustomer(t *testing.T) {
	expectedResult := &model.Customer{
		ID:        1,
		FirstName: "Jeremiah",
		LastName:  "Ferdinand",
		Password:  "123",
		Age:       26,
	}

	query := map[string]interface{}{
		"column":    "tbl_id",
		"separator": "=",
		"value":     1,
	}

	customerRepository.Mock.On("Show", query).Return(expectedResult, nil)

	result := customerService.ShowCustomer(1)

	assert.NotNil(t, *result)
	assert.Equal(t, *expectedResult, *result)
}
