package domain

import "customer/domain/model"

type ICustomerService interface {
	FetchCustomer() *[]model.Customer
	SaveCustomer(customer *model.Customer) error
}

type ICustomerRepository interface {
	Get() *[]model.Customer
	Insert(customer *model.Customer) error
	Show() *model.Customer
	Delete() bool
}
