// Package client provides client implementations to interact with
// other external services such as thrifty.
package client

import (
	"thrifty-client-go/gen-go/thrifty"
)

// ThriftyClient is the client that communicates with the
// thrifty backend service. This types provides all capabilities
// of interacting with the thrifty backend service.
type ThriftyClient struct {
	Client thrifty.PersonService
}

// Get retrieves a thrifty.Person with the provided identifier.
func (thriftyClient *ThriftyClient) Get(id int32) (r *thrifty.Person, err error) {
	return thriftyClient.Client.Get(id)
}

// Create takes the provided thrifty.Person instance and stores it.
func (thriftyClient *ThriftyClient) Create(person *thrifty.Person) (r int32, err error) {
	return thriftyClient.Client.Create(person)
}

// Remove deletes an existing person with the provided identifier.
func (thriftyClient *ThriftyClient) Remove(id int32) (err error) {
	return thriftyClient.Client.Remove(id)
}
