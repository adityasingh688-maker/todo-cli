package main

import (
	"fmt"
	"os"
)

type Task struct {
	ID   int
	Text string
	Done bool
}

var tasks []Task

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo <add|list|done|delete> [args]")
		return
	}

	switch os.Args[1] {
	case "add":
		addTask(os.Args[2:])
	case "list":
		listTasks()
	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
func addTask(args []string) {
	if len(args) == 0 {
		fmt.Println("please provide a task description")
		return
	}
	text := args[0]
	task := Task{ID: len(tasks) + 1, Text: text}
	tasks = append(tasks, task)
	fmt.Println("added:", text)
}
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("no tasks yet")
		return
	}
	for _, t := range tasks {
		status := " "
		if t.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d. %s\n", status, t.ID, t.Text)
	}
}
