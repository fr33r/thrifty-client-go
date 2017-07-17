package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"thrifty-client-go/client"
	"thrifty-client-go/gen-go/thrifty"
)

type App struct {
	client    *client.ThriftyClient
	address   string
	port      string
	transport *thrift.TTransport
}

func (app *App) Connect() error {
	var trans thrift.TTransport = *app.transport
	return trans.Open()
}

func (app *App) Disconnect() {
	var trans thrift.TTransport = *app.transport
	trans.Close()
}

func NewApp(address string, port string) *App {
	transport, err := thrift.NewTSocket(fmt.Sprintf("%v:%v", address, port))
	transportFactory := thrift.NewTTransportFactory()
	trans, err := transportFactory.GetTransport(transport)

	if err != nil {
		panic(err)
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	personServiceClient := thrifty.NewPersonServiceClientFactory(trans, protocolFactory)

	return &App{
		client:    &client.ThriftyClient{Client: personServiceClient},
		address:   address,
		port:      port,
		transport: &trans,
	}
}

func main() {
	app := NewApp("localhost", "9090")

	defer app.Disconnect()

	if err := app.Connect(); err != nil {
		panic(err)
	}

	identifier, err :=
		app.client.Create(&thrifty.Person{GivenName: "Jon", Surname: "Freer", Age: 25})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Person was created with identifier %d.", identifier)
}
