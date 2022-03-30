package repository

import "github.com/gba-3/sample-todo/domain/entity"

type TaskRepository interface {
	GetAll() []entity.Task
}

type taskRepository struct {
}
