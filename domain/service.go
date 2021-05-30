package domain

import "customer/domain/model"

type CustomerService interface {
	FetchCustomer() *[]model.Customer
	SaveCustomer(customer *model.Customer) error
}
