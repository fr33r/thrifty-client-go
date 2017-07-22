package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"thrifty-client-go/client"
	"thrifty-client-go/gen-go/thrifty"
)

// App represents the application as a whole. An instance of this
// struct is bound to a singular IP and port upon instantiation.
// Please see the NewApp constructor function for more details.
type App struct {
	client    *client.ThriftyClient
	address   string
	port      string
	transport *thrift.TTransport
}

// Connect establishes a connection to the IP address and port number
// associated with the calling App instance.
func (app *App) Connect() error {
	var trans thrift.TTransport = *app.transport
	return trans.Open()
}

// Disconnect closes a connection to the IP address and port number
// associated with the calling App instance.
func (app *App) Disconnect() {
	var trans thrift.TTransport = *app.transport
	trans.Close()
}

// NewApp serves as the constructor function for App instances. An IP
// address and port are provided and are required to instantiate each
// App instance.
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
