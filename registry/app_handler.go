package registry

import "github.com/gba-3/sample-todo/handler"

type AppHandler struct {
	Th handler.TaskHandler
	Uh handler.UserHandler
}

func NewAppHandler(th handler.TaskHandler, uh handler.UserHandler) *AppHandler {
	return &AppHandler{th, uh}
}
