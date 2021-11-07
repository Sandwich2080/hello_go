package main

import "fmt"

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
}

// dlv-dap installing guide
// https://github.com/golang/vscode-go/blob/master/docs/debugging.md#updating-dlv-dap
