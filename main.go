package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo <add|list|done|delete> [args]")
		return
	}

	loadTasks()
	switch os.Args[1] {
	case "add":
		addTask(os.Args[2:])
	case "list":
		listTasks()
	case "done":
		doneTask(os.Args[2:])
	case "delete":
		deleteTask(os.Args[2:])
	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
