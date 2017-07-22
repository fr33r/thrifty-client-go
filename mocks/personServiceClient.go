package mocks

import (
	"testify/mock"
	"thrifty-client-go/gen-go/thrifty"
)

type PersonServiceClient struct {
	mock.Mock
}

func (mock *PersonServiceClient) Get(id int32) (*thrifty.Person, error) {
	arguments := mock.Called(id)
	return arguments.Get(0).(*thrifty.Person), arguments.Error(1)
}

func (mock *PersonServiceClient) Create(person *thrifty.Person) (int32, error) {
	arguments := mock.Called(person)
	return int32(arguments.Int(0)), arguments.Error(1)
}

func (mock *PersonServiceClient) Remove(id int32) error {
	arguments := mock.Called(id)
	return arguments.Error(0)
}
