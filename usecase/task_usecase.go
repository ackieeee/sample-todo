package usecase

import (
	"context"

	"github.com/gba-3/sample-todo/domain/entity"
	"github.com/gba-3/sample-todo/domain/repository"
)

type TaskUsecase interface {
	GetAll(ctx context.Context) (entity.Tasks, error)
}

type taskUsecase struct {
	tr repository.TaskRepository
}

func NewTaskUsecase(tr repository.TaskRepository) TaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) GetAll(ctx context.Context) (entity.Tasks, error) {
	tasks, err := tu.tr.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
