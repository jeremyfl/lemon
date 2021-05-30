package domain

import "customer/domain/model"

type CustomerRepository interface {
	Get() *[]model.Customer
	Insert(customer *model.Customer) error
	Show() *model.Customer
	Delete() bool
}
