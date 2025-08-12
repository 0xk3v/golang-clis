package main

import (
	"fmt"
	"os"

	"github.com/0xk3v/golang-clis/todo"
)

const todoFileName = ".todo.json"

func main() {
	// Define a list item
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
