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

// TestComplete tests the Complete method of List type
func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "Drink tea!"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Error("New task should not be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Error("New task should be completed")
	}
}

// TestDelete tests the Delete method of List type
func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"Task one",
		"Task two",
		"Task three",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %d, got %d instead.", 2, len(l))
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length %q, got %q instead.", tasks[0], l[0].Task)
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Task)
	}
}
