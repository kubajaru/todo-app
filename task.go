package main

import "fmt"

type Task struct {
	Title  string
	IsDone bool
	Id     string
}

func (task Task) Print_this() {
	fmt.Println("Task: " + task.Title)
}
