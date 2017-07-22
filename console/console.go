// Package console provides functions and types that facilitate
// interaction with the console.
package console

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Console represents the terminal that receives the
// standard input and outputs to the standard output.
type Console struct {
	scanner *bufio.Scanner
}

// NewConsole is a constructor function for Console instances.
func NewConsole() *Console {
	return &Console{scanner: bufio.NewScanner(os.Stdin)}
}

// Print takes the provided string and prints it to standard output.
func (console *Console) Print(text string) {
	fmt.Print(text)
}

// Print takes the provided string and appends a newline prior to
// printing the string to the standard output.
func (console *Console) Println(text string) {
	fmt.Println(text)
}

// PromptForString prints the provided string to standard output and
// subsequently waits for the next line of input. Returns the received
// input as a string.
func (console *Console) PromptForString(prompt string) (string, error) {
	console.Print(prompt)
	if console.scanner.Scan() {
		return console.scanner.Text(), nil
	}
	return "", errors.New(fmt.Sprintf("An issue occurred when parsing input for prompt '%s'.", prompt))
}

// PromptForInt32 prints the provided string to the standard output and
// subsequently waits for the next line of input. Returns the received
// input as an 32 bit integer.
func (console *Console) PromptForInt32(prompt string) (int32, error) {
	console.Print(prompt)
	if console.scanner.Scan() {
		if integer, err := strconv.Atoi(console.scanner.Text()); err != nil {
			return int32(0), err
		} else {
			return int32(integer), nil
		}
	}
	return int32(0), errors.New(fmt.Sprintf("An issue occurred when parsing input for prompt '%s'.", prompt))
}

// PromptForInt16 prints the provided string to the standard output and
// subsequently waits for the next line of input. Returns the received
// input as an 16 bit integer.
func (console *Console) PromptForInt16(prompt string) (int16, error) {
	console.Print(prompt)
	if console.scanner.Scan() {
		if integer, err := strconv.Atoi(console.scanner.Text()); err != nil {
			return int16(0), err
		} else {
			return int16(integer), nil
		}
	}
	return int16(0), errors.New(fmt.Sprintf("An issue occurred when parsing input for prompt '%s'.", prompt))
}

// PromptForInt8 prints the provided string to the standard output and
// subsequently waits for the next line of input. Returns the received
// input as an 8 bit integer.
func (console *Console) PromptForInt8(prompt string) (int8, error) {
	console.Print(prompt)
	if console.scanner.Scan() {
		if integer, err := strconv.Atoi(console.scanner.Text()); err != nil {
			return int8(0), err
		} else {
			return int8(integer), nil
		}
	}
	return int8(0), errors.New(fmt.Sprintf("An issue occurred when parsing input for prompt '%s'.", prompt))
}
