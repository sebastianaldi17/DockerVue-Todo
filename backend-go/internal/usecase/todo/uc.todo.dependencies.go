package todo

import entityTodo "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/entity/todo"

// go:generate mockgen -build_flags=-mod=mod -source=uc.todo.dependencies.go -package=todo -destination=uc.dependencies.mock_test.go
type todoResource interface {
	AddTodo(todo entityTodo.Todo) error
	DeleteTodo(id int64) error
	GetTodos() ([]entityTodo.Todo, error)
	GetTodoByID(id int64) (entityTodo.Todo, error)
	UpdateTodo(id int64, req entityTodo.UpdateRequest) (int64, error)
}
