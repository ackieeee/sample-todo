package registry

import (
	"database/sql"

	"github.com/gba-3/sample-todo/handler"
	"github.com/gba-3/sample-todo/registry/container"
)

type Registory struct {
	c container.Container
}

func NewRegistory() *Registory {
	return &Registory{
		c: container.Container{},
	}
}

func (r *Registory) Regist(db *sql.DB) *AppHandler {
	return NewAppHandler(
		handler.NewTaskHandler(
			r.c.GetTaskUsecase(
				r.c.GetTaskRepository(db),
			),
		),
		handler.NewUserHandler(
			r.c.GetUserUsecase(
				r.c.GetUserRepository(db),
			),
		),
	)
}
