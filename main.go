package main

import (
	"flag"
	"fmt"
	"strings"
	console2 "thrifty-client-go/console"
	"thrifty-client-go/gen-go/thrifty"
)

func main() {

	host := flag.String("host", "localhost", "the IP address of the host to connect to.")
	port := flag.String("port", "9090", "the port of the host to connect to.")
	flag.Parse()

	app := NewApp(*host, *port)

	defer app.Disconnect()

	if err := app.Connect(); err != nil {
		panic(err)
	}

	console := console2.NewConsole()
	for command, err := console.PromptForString("Enter Command: "); strings.Compare(command, "quit") != 0 && err == nil; {
		if strings.Compare(command, "create") == 0 {
			givenName, err := console.PromptForString("Given name: ")

			if err != nil {
				panic(err)
			}

			surname, err := console.PromptForString("Surname: ")

			if err != nil {
				panic(err)
			}

			age, err := console.PromptForInt8("Age: ")

			if err != nil {
				panic(err)
			}

			identifier, err :=
				app.client.Create(&thrifty.Person{GivenName: givenName, Surname: surname, Age: age})

			if err != nil {
				panic(err)
			}

			fmt.Printf("Person was created with identifier %d.\n", identifier)
		} else if strings.Compare(command, "get") == 0 {
			id, err := console.PromptForInt32("Id: ")

			if err != nil {
				panic(err)
			}

			person, err := app.client.Get(int32(id))

			if err != nil {
				panic(err)
			}

			fmt.Println(person)
		} else if strings.Compare(command, "remove") == 0 {
			id, err := console.PromptForInt32("Id: ")

			if err != nil {
				panic(err)
			}

			err = app.client.Remove(id)

			if err != nil {
				panic(err)
			}

			fmt.Printf("Person with id %d has been removed.\n", id)
		} else {
			console.Println("Command not recognized. Please use one of the following commands: [get, create, remove]")
		}
		command, err = console.PromptForString("Enter Command: ")
	}
	console.Println("See ya later!")
}
