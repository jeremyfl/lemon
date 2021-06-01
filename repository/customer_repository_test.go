package repository

import (
	"context"
	"customer/domain/model"
	"customer/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

var (
	mockDb = mocks.Database{Mock: mock.Mock{}}

	customerRepository = CustomerRepositoryImpl{
		Ctx: context.Background(),
		DB:  &mockDb,
	}
)

func TestCustomerRepositoryImpl_Delete(t *testing.T) {
}

func TestCustomerRepositoryImpl_Get(t *testing.T) {
	var customers []model.Customer

	mockDb.Mock.On("Select", &customers).Return(nil)
	result := customerRepository.Get()

	assert.Nil(t, *result)
}

func TestCustomerRepositoryImpl_Insert(t *testing.T) {
	payload := &model.Customer{
		ID:        0,
		FirstName: "",
		LastName:  "",
		Password:  "",
		Age:       0,
	}

	mockDb.On("Insert", payload).Return(payload.ID, nil)

	result := customerRepository.Insert(payload)

	assert.Nil(t, result)
}

func TestCustomerRepositoryImpl_Show(t *testing.T) {
	var cs []model.Customer

	query := map[string]interface{}{
		"column":    "tbl_id",
		"separator": "=",
		"value":     1,
	}

	mockDb.Mock.On("Select", &cs, query).Return(nil)

	result, _ := customerRepository.Show(query)

	assert.NotNil(t, result)
}
