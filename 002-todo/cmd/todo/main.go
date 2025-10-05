package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/0xk3v/golang-clis/todo"
)

const todoFileName = ".todo.json"

func main() {
	// Parsing CMD arguments
	task := flag.String("task", "", "Task to be included in the ToDo List")
	list := flag.Bool("list", false, "Lists all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	// Define a list item
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(l)

	case *task != "":
		l.Add(*task)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	case *complete > 0:
		// Complete the given task.
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
