package todo_test

import (
	"testing"

	"github.com/0xk3v/golang-clis/todo"
)

// TestAdd tests the Add method of the List type
func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "Drink tea!"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}
