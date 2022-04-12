package registry

import "github.com/gba-3/sample-todo/handler"

type AppHandler struct {
	Th handler.TaskHandler
}

func NewAppHandler(th handler.TaskHandler) *AppHandler {
	return &AppHandler{th}
}
