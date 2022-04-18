package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gba-3/sample-todo/domain/entity"
	"github.com/gba-3/sample-todo/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	Find(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, name, email, password string) error
	GetAll(ctx context.Context) (entity.Users, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Find(ctx context.Context, email string) (*entity.User, error) {
	u, err := models.Users(
		qm.Where("email=?", email),
	).One(ctx, ur.db)
	if err != nil {
		return nil, err
	}
	user := entity.User{
		ID:       int(u.ID),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	return &user, nil
}

func (ur *userRepository) Create(ctx context.Context, name, email, password string) error {
	if name == "" {
		return errors.New("arguments error: name is empty.")
	}
	if email == "" {
		return errors.New("arguments error: email is empty.")
	}
	if password == "" {
		return errors.New("arguments error: password is empty.")
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return user.Insert(ctx, ur.db, boil.Infer())
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
