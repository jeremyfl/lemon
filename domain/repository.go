package domain

import "customer/domain/model"

type CustomerService interface {
	FetchCustomer() *[]model.Customer
	SaveCustomer(customer *model.Customer) error
}

type CustomerRepository interface {
	Get() *[]model.Customer
	Insert(customer *model.Customer) error
	Show() *model.Customer
	Delete() bool
}
