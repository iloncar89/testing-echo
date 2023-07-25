package service

import (
	"testing-echo/domain"
	"testing-echo/model"
)

type TestService interface {
	CalculateFibonacci(number int64) int64
	CreateGetDeletePersonTestCase(person model.Person) (model.Person, error)
	CreateGetDeletePersonORMTestCase(person domain.Person) (domain.Person, error)
}
