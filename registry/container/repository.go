package container

import (
	"database/sql"

	"github.com/gba-3/sample-todo/domain/repository"
)

func (c Container) GetTaskRepository(db *sql.DB) repository.TaskRepository {
	return repository.NewTaskRepository(db)
}

func (c Container) GetUserRepository(db *sql.DB) repository.UserRepository {
	return repository.NewUserRepository(db)
}
