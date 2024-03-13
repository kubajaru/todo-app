package task

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SqliteRepository struct {
	db *sql.DB
}

func NewSqliteRepository(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (r *SqliteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL UNIQUE,
        is_done BOOLEAN NOT NULL,
    );
	`

	_, err := r.db.Exec(query)

	return err
}

func (r *SqliteRepository) Create(task Task) (*Task, error) {
	res, err := r.db.Exec("INSERT INTO tasks(title, is_done) values(?,?,?)", task.title, task.isDone)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				slog.Error("Task: " + task.title + " aleardy exists.")
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	task.id = id

	return &task, nil
}
