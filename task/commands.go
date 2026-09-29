package task

import (
	"fmt"
	"strconv"
)

func AddTask(args []string) {
	if len(args) == 0 {
		fmt.Println("please provide a task description")
		return
	}
	text := args[0]
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	task := Task{ID: maxID + 1, Text: text}
	tasks = append(tasks, task)
	saveTasks()
	fmt.Println("added:", text)
}

func ListTasks() {
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

func DoneTask(args []string) {
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

func DeleteTask(args []string) {
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
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Println("deleted task", id)
			return
		}
	}
	fmt.Println("task not found")
}
