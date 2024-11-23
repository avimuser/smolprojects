package main

import (
	"os"
	"strconv"

	"github.com/vinaykandagatla/backend-projects/todo/cmd"
	"github.com/vinaykandagatla/backend-projects/todo/internal/task"
)

func main() {
	arguments := os.Args

	if len(arguments) < 2 {
		cmd.ListTasks(-1)
		os.Exit(0)
	}

	switch arguments[1] {
	case "add":
		if len(arguments) < 3 {
			cmd.Help()
			os.Exit(1)
		}

		cmd.AddTask(arguments[2])
	case "update":
		if len(arguments) < 4 {
			cmd.Help()
			os.Exit(1)
		}

		id, _ := strconv.Atoi(arguments[2])
		cmd.UpdateTask(id, arguments[3])
	case "delete":
		if len(arguments) < 3 {
			cmd.Help()
			os.Exit(1)
		}

		id, _ := strconv.Atoi(arguments[2])
		cmd.DeleteTask(id)
	case "mark":
		if len(arguments) < 4 {
			cmd.Help()
			os.Exit(0)
		}

		index, _ := strconv.Atoi(arguments[2])
		switch arguments[3] {
		case "todo":
			cmd.MarkTask(index, task.Incomplete)
		case "in-progress":
			cmd.MarkTask(index, task.Started)
		case "done":
			cmd.MarkTask(index, task.Completed)
		default:
			cmd.Help()
		}
	case "list":
		if len(arguments) < 3 {
			cmd.ListTasks(-1)
			os.Exit(0)
		}

		switch arguments[2] {
		case "todo":
			cmd.ListTasks(task.Incomplete)
		case "in-progress":
			cmd.ListTasks(task.Started)
		case "done":
			cmd.ListTasks(task.Completed)
		default:
			cmd.Help()
		}
	case "help":
		cmd.Help()
	default:
		cmd.Help()
	}
}
