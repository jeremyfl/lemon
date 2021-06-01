package service

import (
	"customer/domain"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	Repository domain.CustomerRepository
}

func (a AuthServiceImpl) Login(payload ...string) error {
	q1 := map[string]interface{}{
		"column":    "email",
		"separator": "=",
		"value":     payload[0],
	}

	customer, err := a.Repository.Show(q1)
	if err != nil {
		return err
	}

	if customer.Email == "" {
		return errors.New("unauthorized")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(payload[1])); err != nil {
		fmt.Println(err.Error())
		return errors.New("wrong_password")
	}

	return nil
}
