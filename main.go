package main

import (
	"database/sql"
	"log/slog"
	"os"

	"example/todo-app/task"
)

const fileName = "sqlite.db"

func main() {
	os.Remove(fileName)

	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		slog.Error(err.Error())
	}
	taskRepository := task.NewSqliteRepository(db)

	task := task.NewTask("test", false)

	taskRepository.Migrate()

	taskRepository.Create(*task)
}
