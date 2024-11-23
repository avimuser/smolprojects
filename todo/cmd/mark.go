package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/vinaykandagatla/backend-projects/todo/internal/task"
)

func MarkTask(id int, status task.TaskStatus) {
	index := id - 1
	raw, err := os.ReadFile("tasks.json")

	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(err)
		}
		os.Exit(0)
	}

	tasks := []task.Task{}
	json.Unmarshal(raw, &tasks)

	if index >= len(tasks) || index < 0 {
		fmt.Println("Task with that ID does not exist")
		os.Exit(1)
	}

	tasks[index].Status = status
	tasks[index].UpdatedAt = time.Now()

	payload, _ := json.Marshal(tasks)
	os.WriteFile("tasks.json", payload, 0644)
}
