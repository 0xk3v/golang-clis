// Package todo
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represends a list of ToDo items
type List []item

// Add creates a new ToDo item
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// Complete marks a given task as 'Done'
func (l *List) Complete(id int) error {
	ls := *l

	if id <= 0 || id > len(ls) {
		return fmt.Errorf("item %d does not exist", id)
	}

	ls[id-1].Done = true
	ls[id-1].CompletedAt = time.Now()

	return nil
}

// Delete removes an item from the list of ToDos
func (l *List) Delete(id int) error {
	ls := *l

	if id <= 0 || id > len(ls) {
		return fmt.Errorf("item %d does not exist", id)
	}

	*l = append(ls[:id-1], ls[id:]...)

	return nil
}

// Save method encodes the List as JSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

// String prints out a formatted list
// Implements the fmt.Stringer interface
func (l *List) String() string {
	formatted := ""

	for index, item := range *l {

		prefix := "  "
		if item.Done {
			prefix = "X "
		}

		formatted += fmt.Sprintf("%s%d: %s\n", prefix, index+1, item.Task)

	}

	return formatted
}
