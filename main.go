package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/gba-3/sample-todo/models"
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
	ctx := context.Background()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("call")
		tasks, err := models.Tasks().All(ctx, db)
		if err != nil {
			log.Println(err.Error())
		}
		for _, task := range tasks {
			log.Println(task.Title)
		}
		w.Write([]byte("success"))
	})
	http.ListenAndServe(":3000", r)
}
