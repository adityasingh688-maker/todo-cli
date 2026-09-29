package task

type Task struct {
	ID   int
	Text string
	Done bool
}

const filename = "todos.json"

var tasks []Task
