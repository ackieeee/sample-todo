package handler

import (
	"context"
	"net/http"

	"github.com/gba-3/sample-todo/usecase"
)

type userHandler struct {
	uu usecase.UserUsecase
}

type UserHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{uu}
}

func (uh *userHandler) GetAll(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	ctx := context.Background()
	users, err := uh.uu.GetAll(ctx)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, users, nil
}
