package main

import "fmt"

func main() {
	fmt.Println("Heelo world!")
	task := new(Task)
	task.Id = "11111"
	task.IsDone = true
	task.Title = "Clean the room"
	task.Print_this()
}
