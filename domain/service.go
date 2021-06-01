package domain

import "customer/domain/model"

type CustomerService interface {
	FetchCustomer() *[]model.Customer
	ShowCustomer(id int) *model.Customer
	SaveCustomer(customer *model.Customer) error
}
