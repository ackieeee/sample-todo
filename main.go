package main

import (
	"database/sql"
	"net/http"

	"github.com/gba-3/sample-todo/registry"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := sql.Open("mysql", "todo:todo@tcp(db)/todo?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	ah := registry.NewRegistory().Regist(db)
	r.Get("/tasks", ah.Th.GetAll)
	r.Post("/tasks/add", ah.Th.AddTask)
	r.Post("/tasks/status/update", ah.Th.ChangeStatus)
	http.ListenAndServe(":3000", r)
}
