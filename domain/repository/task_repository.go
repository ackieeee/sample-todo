package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/gba-3/sample-todo/domain/entity"
	"github.com/gba-3/sample-todo/models"
)

type TaskRepository interface {
	GetAll(ctx context.Context) (entity.Tasks, error)
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAll(ctx context.Context) (entity.Tasks, error) {
	es := entity.Tasks{}
	tasks, err := models.Tasks().All(ctx, tr.db)
	if err != nil {
		log.Panic(err.Error())
		return nil, err
	}
	for _, task := range tasks {
		e := entity.Task{
			Title:       task.Title,
			Description: task.Description.String,
			Date:        task.Date.Time.String(),
			Status:      task.Status.Bool,
		}
		es = append(es, e)
	}
	return es, nil
}
