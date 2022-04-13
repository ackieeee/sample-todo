package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gba-3/sample-todo/domain/entity"
	"github.com/gba-3/sample-todo/usecase"
)

type taskhandler struct {
	tu usecase.TaskUsecase
}

type TaskHandler interface {
	ChangeStatus(w http.ResponseWriter, r *http.Request)
	AddTask(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

func NewTaskHandler(tu usecase.TaskUsecase) TaskHandler {
	return &taskhandler{tu}
}

func (th *taskhandler) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	task := entity.Task{}
	if err := json.Unmarshal(body, &task); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if task.ID == 0 {
		w.Write([]byte("task id is empty"))
		return
	}
	err = th.tu.ChangeStatus(ctx, task.ID, task.Status)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("success: ChangeStatus"))
}

func (th *taskhandler) AddTask(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	task := entity.Task{}
	if err := json.Unmarshal(body, &task); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if err := th.tu.AddTask(ctx, task.Title, task.Description, task.Date); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("success: add task."))

}

func (th *taskhandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	tasks, err := th.tu.GetAll(ctx)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte("failed get all tasks."))
		return
	}
	jsonData, err := tasks.ToJson()
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte("failed convert to json."))
		return
	}
	w.Write(jsonData)
}
