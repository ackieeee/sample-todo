package container

import (
	"github.com/gba-3/sample-todo/domain/repository"
	"github.com/gba-3/sample-todo/usecase"
)

func (c Container) GetTaskUsecase(tr repository.TaskRepository) usecase.TaskUsecase {
	return usecase.NewTaskUsecase(tr)
}

func (c Container) GetUserUsecase(ur repository.UserRepository) usecase.UserUsecase {
	return usecase.NewUserUsecase(ur)
}
