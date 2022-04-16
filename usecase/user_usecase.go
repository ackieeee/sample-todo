package usecase

import (
	"context"

	"github.com/gba-3/sample-todo/domain/entity"
	"github.com/gba-3/sample-todo/domain/repository"
)

type userUsecase struct {
	ur repository.UserRepository
}

type UserUsecase interface {
	GetAll(ctx context.Context) (entity.Users, error)
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		ur,
	}
}

func (uu *userUsecase) GetAll(ctx context.Context) (entity.Users, error) {
	return uu.ur.GetAll(ctx)
}
