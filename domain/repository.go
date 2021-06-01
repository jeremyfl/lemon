package domain

import "customer/domain/model"

type CustomerRepository interface {
	Get() *[]model.Customer
	Insert(customer *model.Customer) error
	Show(query ...map[string]interface{}) (*model.Customer, error)
	Delete() bool
}
