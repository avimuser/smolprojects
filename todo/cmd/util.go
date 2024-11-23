package cmd

import "fmt"

func Help() {
	fmt.Println("Todo help:")
	fmt.Println("")
	fmt.Println("Available commands:")
	fmt.Println("  add <task>                           Add a new task")
	fmt.Println("  list <done|in-progress|todo>         List all tasks")
	fmt.Println("  update <id> <task>                   Update an existing task")
	fmt.Println("  delete <id>                          Delete an existing task")
	fmt.Println("  mark <id> <done|in-progress|todo>    Delete an existing task")
	fmt.Println("  help                                 Show this help message")
}
