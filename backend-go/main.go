package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	todoHTTP "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/handler/http/todo"
	todoResource "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/resource/todo"
	todoUsecase "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/usecase/todo"
)

type handlers struct {
	todoHTTP *todoHTTP.Handler
}

func main() {
	connStr := "postgres://root:root@postgres/docker-todo-db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
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
	r.Use(middleware.Logger)

	r.HandleFunc("/", h.todoHTTP.Hello)

	http.ListenAndServe(":3000", r)
}
