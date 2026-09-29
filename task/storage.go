package task

import (
	"encoding/json"
	"fmt"
	"os"
)

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

func LoadTasks() {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println(err)
	}
}
