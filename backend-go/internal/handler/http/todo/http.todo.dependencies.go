package todo

import entityTodo "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/entity/todo"

//go:generate mockgen -build_flags=-mod=mod -source=http.todo.dependencies.go -package=todo -destination=http.dependencies.mock_test.go
type todoUsecase interface {
	AddTodo(todo entityTodo.Todo) error
	DeleteTodo(id int64) error
	GetTodos() ([]entityTodo.Todo, error)
	GetTodoByID(id int64) (entityTodo.Todo, error)
	UpdateTodo(req entityTodo.UpdateRequest, id int64) (int64, error)
}
