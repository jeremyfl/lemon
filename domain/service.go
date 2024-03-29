package domain

import "customer/domain/model"

type Services struct {
	AuthService
	CustomerService
}

type AuthService interface {
	Login(payload ...string) error
}

type CustomerService interface {
	FetchCustomer() *[]model.Customer
	ShowCustomer(id int) *model.Customer
	SaveCustomer(customer *model.Customer) error
	UpdateCustomer(customer *model.Customer, id int64) error
	DeleteCustomer(id int64) error
}
