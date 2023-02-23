package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	todoHTTP "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/handler/http/todo"
	todoResource "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/resource/todo"
	todoUsecase "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/usecase/todo"
)

type handlers struct {
	todoHTTP *todoHTTP.Handler
}

func main() {
	connStr := "postgres://root:root@127.0.0.1/docker-todo-db?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	todoRes := todoResource.New(db)

	todoUC := todoUsecase.New(todoRes)

	var h handlers

	todoHTTP := todoHTTP.New(todoUC)
	h.todoHTTP = todoHTTP

	setupRouter(h)
}

func setupRouter(h handlers) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", h.todoHTTP.Hello)
	r.Get("/todos", h.todoHTTP.GetTodos)

	r.Route("/todo", func(r chi.Router) {
		r.Post("/", h.todoHTTP.AddTodo)

		r.Route("/{id}", func(r chi.Router) {
			r.Delete("/", h.todoHTTP.DeleteTodo)
			r.Put("/", h.todoHTTP.UpdateTodo)
		})
	})

	log.Println("Listening to :3000")
	http.ListenAndServe(":3000", r)
}
