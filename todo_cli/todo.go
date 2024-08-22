package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

// add todo function
func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

// helper function to validate the index of the given operation
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// delete and item from the list
func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...) // index ke pehle ke elem and uske baad ke elems ko join kr de rha h us index ko chod kr

	return nil
}

// toggle function to change the status of completed or not
func (todos *Todos) toggle(index int) {
	if index < 0 || index >= len(*todos) {
		fmt.Println("Invalid index")
		return
	}

	t := &(*todos)[index] // Get a pointer to the struct in the slice
	t.Completed = !t.Completed

	if t.Completed {
		now := time.Now()
		t.CompletedAt = &now
	} else {
		t.CompletedAt = nil
	}
}

// function to edit the tasks/todos
func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil

}

// function to print all the todos in terminal using table package
func (todos *Todos) print() {
	table := table.New(os.Stdout)

	// configure the table to print result ina desired way
	table.SetRowLines(true) // no horizontal lines
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	// loop through all the todos and render it on terminal
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}

		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()

}
