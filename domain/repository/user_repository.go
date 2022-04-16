package repository

import (
	"context"
	"database/sql"

	"github.com/gba-3/sample-todo/domain/entity"
	"github.com/gba-3/sample-todo/models"
)

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	GetAll(ctx context.Context) (entity.Users, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetAll(ctx context.Context) (entity.Users, error) {
	users, err := models.Users().All(ctx, ur.db)
	if err != nil {
		return nil, err
	}

	us := entity.Users{}
	for _, user := range users {
		u := entity.User{
			ID:       int(user.ID),
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
		us = append(us, u)
	}
	return us, nil
}
