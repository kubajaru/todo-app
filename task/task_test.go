package task

import "testing"

func TestNewTask(t *testing.T) {
	got := NewTask("test", true)
	want := "test"
	wantIsDone := true

	if got.title != want {
		t.Errorf("got %q, wanted %q", got.title, want)
	}

	if got.isDone != wantIsDone {
		t.Errorf("got %t, wanted %t", got.isDone, wantIsDone)
	}
}
