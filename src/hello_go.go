package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

type Book struct {
	id      int
	title   string
	author  string
	subject string
}

func main() {

	var name = "Mr Jackson"
	passcode := "123"
	var age int = 23
	fmt.Println("name: "+name+",passcode:"+passcode+",age:", age)

	var book1 = Book{id: 1, title: "Harry Portter", author: "Unknown", subject: "unknown"}
	fmt.Println(book1)

	// Get a greeting message and print it.
	// message := greetings.Hello("Programming lovers")
	// fmt.Println(message)

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(messages)
}

// dlv-dap installing guide
// https://github.com/golang/vscode-go/blob/master/docs/debugging.md#updating-dlv-dap

// Guide for new Go programmer
// https://golang.org/doc/tutorial/random-greeting
