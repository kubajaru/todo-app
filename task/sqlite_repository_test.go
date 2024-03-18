package task

import (
	"database/sql"
	"testing"
)

func TestMigrate(t *testing.T) {

	db, err := sql.Open("sqlite3", "/tmp/sqlite.db")
	if err != nil {
		t.Fatalf(err.Error())
	}

	sql := NewSqliteRepository(db)

	want := NewTask("test", false)

	got, err := sql.Create(*want)

	if got.title != want.title {
		t.Errorf("got %q, wanted %q", got.title, want.title)
	}

	if got.isDone != want.isDone {
		t.Errorf("got %t, wanted %t", got.isDone, want.isDone)
	}

	t.Cleanup(func() {
		db.Close()
	})
}
