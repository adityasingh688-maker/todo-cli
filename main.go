package main

import (
	"github.com/adityasingh688-maker/todo-cli/task"
	"github.com/gin-gonic/gin"
)

func main() {
	task.LoadTasks()

	r := gin.Default()

	r.GET("/tasks", func(c *gin.Context) {
		tasks := task.ListTasks()
		c.JSON(200, tasks)
	})

	r.POST("/tasks", func(c *gin.Context) {
		var body struct {
			Text string `json:"text"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		newTask, err := task.AddTask(body.Text)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, newTask)
	})

	r.Run(":8080")
}
