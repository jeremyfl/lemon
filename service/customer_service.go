package service

import (
	"customer/domain"
	"customer/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type CustomerServiceImpl struct {
	Repository domain.CustomerRepository
}

func (cs CustomerServiceImpl) FetchCustomer() *[]model.Customer {
	return cs.Repository.Get()
}

func (cs CustomerServiceImpl) ShowCustomer(id int) *model.Customer {
	query := map[string]interface{}{
		"column":    "tbl_id",
		"separator": "=",
		"value":     id,
	}

	customer, _ := cs.Repository.Show(query)

	return customer
}

func (cs CustomerServiceImpl) SaveCustomer(customer *model.Customer) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashedPassword := string(hashPassword)
	customer.Password = hashedPassword

	return cs.Repository.Insert(customer)
}
