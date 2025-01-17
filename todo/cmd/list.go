package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/avimuser/smolprojects/todo/internal/task"
)

func ListTasks(status task.TaskStatus) {
	raw, err := os.ReadFile("tasks.json")
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(err)
		}
		os.Exit(0)
	}

	tasks := []task.Task{}
	json.Unmarshal(raw, &tasks)
	for i := 0; i < len(tasks); i++ {
		if status == -1 || tasks[i].Status == status {
			symbols := []string{"❌", "🔃", "✅"}
			symbol := symbols[tasks[i].Status]
			fmt.Println(symbol, i+1, "-", tasks[i].Desc)
		}
	}
}
