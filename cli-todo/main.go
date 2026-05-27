package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const filename = "todos.json"

func saveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func loadTodos() ([]Todo, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil // first run — no file yet
		}
		return nil, err
	}
	var todos []Todo
	err = json.Unmarshal(data, &todos)
	return todos, err
}

func printTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("No todos yet.")
		return
	}
	for _, todo := range todos {
		status := "[ ]"
		if todo.Completed {
			status = "[✓]"
		}
		fmt.Printf("%d. %s %s\n", todo.ID, status, todo.Title)
	}
}

func main() {
	// load existing todos
	todos, err := loadTodos()
	if err != nil {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}

	// add two items
	todos = append(todos, Todo{ID: len(todos) + 1, Title: "Buy groceries"})
	todos = append(todos, Todo{ID: len(todos) + 1, Title: "Write tests"})

	// mark first done
	todos[0].Completed = true

	// print
	printTodos(todos)

	// save back to disk
	if err := saveTodos(todos); err != nil {
		fmt.Println("Error saving todos:", err)
		os.Exit(1)
	}

	fmt.Println("\n✓ Saved to", filename)
}