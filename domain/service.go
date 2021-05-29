package domain

type CustomerServiceFunction interface {
	AuthenticateUser() bool
}
