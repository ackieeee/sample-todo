package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gba-3/sample-todo/domain/entity"
	"github.com/gba-3/sample-todo/usecase"
)

type taskhandler struct {
	tu usecase.TaskUsecase
}

type TaskHandler interface {
	ChangeStatus(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
	AddTask(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
	GetAll(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
}

func NewTaskHandler(tu usecase.TaskUsecase) TaskHandler {
	return &taskhandler{tu}
}

func (th *taskhandler) ChangeStatus(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	task := entity.Task{}
	if err := json.Unmarshal(body, &task); err != nil {
		return http.StatusBadRequest, nil, err
	}
	if task.ID == 0 {
		return http.StatusBadRequest, nil, errors.New("task id is empty.")
	}
	err = th.tu.ChangeStatus(ctx, task.ID, task.Status)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, "success", nil
}

func (th *taskhandler) AddTask(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	ctx := context.Background()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	task := entity.Task{}
	if err := json.Unmarshal(body, &task); err != nil {
		return http.StatusBadRequest, nil, err
	}
	if err := th.tu.AddTask(ctx, task.Title, task.Description, task.Date); err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, "success", nil
}

func (th *taskhandler) GetAll(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	tasks, err := th.tu.GetAll(r.Context())
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, tasks, nil
}
