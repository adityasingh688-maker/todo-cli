package main

import (
	"fmt"
	"os"

	"github.com/adityasingh688-maker/todo-cli/task"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo <add|list|done|delete> [args]")
		return
	}

	task.LoadTasks()
	switch os.Args[1] {
	case "add":
		task.AddTask(os.Args[2:])
	case "list":
		task.ListTasks()
	case "done":
		task.DoneTask(os.Args[2:])
	case "delete":
		task.DeleteTask(os.Args[2:])
	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
