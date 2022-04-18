package usecase

import (
	"context"
	"errors"

	"github.com/gba-3/sample-todo/domain/entity"
	"github.com/gba-3/sample-todo/domain/repository"
)

type userUsecase struct {
	ur repository.UserRepository
}

type UserUsecase interface {
	Find(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, name string, email string, password string) error
	GetAll(ctx context.Context) (entity.Users, error)
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		ur,
	}
}
func (uu *userUsecase) Find(ctx context.Context, email string) (*entity.User, error) {
	if email == "" {
		return nil, errors.New("arguments error: email is empty")
	}
	return uu.ur.Find(ctx, email)
}
func (uu *userUsecase) Create(ctx context.Context, name string, email string, password string) error {
	return uu.ur.Create(ctx, name, email, password)
}

func (uu *userUsecase) GetAll(ctx context.Context) (entity.Users, error) {
	return uu.ur.GetAll(ctx)
}
