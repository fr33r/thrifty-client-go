package client

import (
	"thrifty-client-go/gen-go/thrifty"
)

type ThriftyClient struct {
	Client *thrifty.PersonServiceClient
}

func (thriftyClient *ThriftyClient) Get(id int32) (r *thrifty.Person, err error) {
	return thriftyClient.Client.Get(id)
}

func (thriftyClient *ThriftyClient) Create(person *thrifty.Person) (r int32, err error) {
	return thriftyClient.Client.Create(person)
}

func (thriftyClient *ThriftyClient) Remove(id int32) (err error) {
	return thriftyClient.Client.Remove(id)
}
