package usecase

import (
	"context"

	"github.com/gba-3/sample-todo/domain/entity"
	"github.com/gba-3/sample-todo/domain/repository"
)

type TaskUsecase interface {
	ChangeStatus(ctx context.Context, id int, status bool) error
	AddTask(ctx context.Context, title string, description string, date string) error
	GetAll(ctx context.Context) (entity.Tasks, error)
}

type taskUsecase struct {
	tr repository.TaskRepository
}

func NewTaskUsecase(tr repository.TaskRepository) TaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) ChangeStatus(ctx context.Context, id int, status bool) error {
	return tu.tr.ChangeStatus(ctx, id, status)
}

func (tu *taskUsecase) AddTask(ctx context.Context, title string, description string, date string) error {
	return tu.tr.AddTask(ctx, title, description, date)
}

func (tu *taskUsecase) GetAll(ctx context.Context) (entity.Tasks, error) {
	tasks, err := tu.tr.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
