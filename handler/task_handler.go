package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gba-3/sample-todo/usecase"
)

type taskhandler struct {
	tu usecase.TaskUsecase
}

type TaskHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

func NewTaskHandler(tu usecase.TaskUsecase) TaskHandler {
	return &taskhandler{tu}
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
