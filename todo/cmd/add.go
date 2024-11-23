package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/vinaykandagatla/backend-projects/todo/internal/task"
)

func AddTask(desc string) {
	raw, err := os.ReadFile("tasks.json")

	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(err)
			os.Exit(0)
		}
	}

	tasks := []task.Task{}
	json.Unmarshal(raw, &tasks)
	tasks = append(tasks, task.Task{Desc: desc, Status: task.Incomplete, CreateAt: time.Now(), UpdatedAt: time.Now()})

	payload, _ := json.Marshal(tasks)
	os.WriteFile("tasks.json", payload, 0644)
}
