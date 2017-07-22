# thrifty-client-go

The main purpose of this repository is to provide a 
working example of how to create a client using [Apache 
Thrift](https://thrift.apache.org) and 
[golang](https://golang.org). This client provides the 
capability to interact with a server application called
[thrifty](https://github.com/freerjm/thrifty).

### Up and Running

Here are few steps that should help you get everything 
setup and running:

1. Clone the server application (called thrifty) that 
this client communicates with [here](https://github.com/freerjm/thrifty).

2. Once in the directory of thrifty, run the following 
`bash` command: `go build && ./thrifty`

3. Clone this repository into your `go` workspace.

4. Once in the directory of `thrifty-client-go`, run the
following `bash` command: `go build && ./thrifty-client-go`

### Usage

Once you have both `thrifty`, and `thrifty-client-go` up
and running, use the following commands with `thrifty-client-go`
to perform supported operations:

- `create` - creates and stores a new person.

- `get` - retrieves an already existing person by specifying
and identifier.

- `remove` - removes an already existing person with the specified
identifier.

### Examples

```text
Enter Command: create
Given name: Jonny
Surname: Testerson
Age: 25
Person was created with identifier 4.
```

```text
Enter Command: get
Id: 4
Person({GivenName:Jonny Surname:Testerson Age:25 Address:Address({Line1: Line2:<nil> City: State: Country:})})
```

```text
Enter Command: remove
Id: 4
Person with id 4 has been removed.
```