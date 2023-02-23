package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	connStr := "postgres://root:root@%s/docker-todo-db?sslmode=disable"
	if _, err := os.Stat("/.dockerenv"); err != nil {
		connStr = fmt.Sprintf(connStr, "127.0.0.1")
	} else {
		connStr = fmt.Sprintf(connStr, "postgres")
	}
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
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", h.todoHTTP.Hello)
	r.Get("/todos", h.todoHTTP.GetTodos)

	r.Route("/todo", func(r chi.Router) {
		r.Post("/", h.todoHTTP.AddTodo)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.todoHTTP.GetTodoByID)
			r.Delete("/", h.todoHTTP.DeleteTodo)
			r.Put("/", h.todoHTTP.UpdateTodo)
		})
	})

	log.Println("Listening to :3000")
	http.ListenAndServe(":3000", r)
}
