package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID   int
	Text string
	Done bool
}

const filename = "todos.json"

var tasks []Task

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
	saveTasks()
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
			status = "X"
		}
		fmt.Printf("[%s] %d. %s\n", status, t.ID, t.Text)
	}
}
func saveTasks() {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("error saving tasks:", err)
		return
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("error writing file:", err)
	}
}
func loadTasks() {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println(err)
	}
}
func doneTask(args []string) {
	if len(args) == 0 {
		fmt.Println("please provide a task ID")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("invalid ID:", args[0])
		return
	}
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			saveTasks()
			fmt.Println("done:", tasks[i].Text)
			return
		}
	}
	fmt.Println("task not found")
}
