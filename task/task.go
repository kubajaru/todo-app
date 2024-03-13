package task

import "fmt"

type Task struct {
	title  string
	isDone bool
	id     int64
}

func (task Task) Print_this() {
	fmt.Println("Task: " + task.title)
}

func NewTask(title string, is_done bool) *Task {
	return &Task{
		title:  title,
		isDone: is_done,
	}
}
