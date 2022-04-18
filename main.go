package main

import (
	"database/sql"
	"net/http"

	"github.com/gba-3/sample-todo/handler"
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
	r.Get("/tasks", handler.JsonHandler(ah.Th.GetAll).ServeHTTP)
	r.Post("/tasks/add", handler.JsonHandler(ah.Th.AddTask).ServeHTTP)
	r.Post("/tasks/status/update", handler.JsonHandler(ah.Th.ChangeStatus).ServeHTTP)
	r.Get("/users", handler.JsonHandler(ah.Uh.GetAll).ServeHTTP)
	r.Post("/users/add", handler.JsonHandler(ah.Uh.Signup).ServeHTTP)
	r.Post("/login", handler.JsonHandler(ah.Uh.Login).ServeHTTP)
	http.ListenAndServe(":3000", r)
}
