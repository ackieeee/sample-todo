package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gba-3/sample-todo/domain/entity"
	"github.com/gba-3/sample-todo/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TaskRepository interface {
	ChangeStatus(ctx context.Context, id int, status bool) error
	AddTask(ctx context.Context, title string, description string, date string) error
	GetAll(ctx context.Context) (entity.Tasks, error)
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) ChangeStatus(ctx context.Context, id int, status bool) error {
	task, err := models.FindTask(ctx, tr.db, uint64(id))
	if err != nil {
		return err
	}
	if task.Status == status {
		return nil
	}
	task.Status = status
	_, err = task.Update(ctx, tr.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil

}

func (tr *taskRepository) AddTask(ctx context.Context, title string, description string, date string) error {
	tx, err := tr.db.Begin()
	if err != nil {
		return err
	}
	d := null.NewString(description, true)
	if description == "" {
		d = null.NewString("", false)
	}

	t, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return err
	}
	task := models.Task{
		Title:       title,
		Description: d,
		Date:        null.NewTime(t, true),
	}
	if err := task.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
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
			ID:          int(task.ID),
			Title:       task.Title,
			Description: task.Description.String,
			Date:        task.Date.Time.String(),
			Status:      task.Status,
		}
		es = append(es, e)
	}
	return es, nil
}
