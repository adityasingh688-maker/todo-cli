package task

import (
	"fmt"
)

func AddTask(text string) (Task, error) {
	if text == "" {
		return Task{}, fmt.Errorf("task description cannot be empty")
	}
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	task := Task{ID: maxID + 1, Text: text}
	tasks = append(tasks, task)
	saveTasks()
	return task, nil
}

func ListTasks() []Task {
	return tasks
}

func DoneTask(id int) (Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			saveTasks()
			return tasks[i], nil
		}
	}
	return Task{}, fmt.Errorf("task not found")
}

func DeleteTask(id int) error {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			return nil
		}
	}
	return fmt.Errorf("task not found")
}
