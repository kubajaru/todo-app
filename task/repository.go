package task

type TaskRepository interface {
	Migrate() error
	Create(task Task) (*Task, error)
	All() ([]Task, error)
	GetByName(name string) (*Task, error)
	Update(id int64, updated Task) (*Task, error)
	Delete(id int64) error
}
