package todo

import entityTodo "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/entity/todo"

func (u *Usecase) GetTodos() ([]entityTodo.Todo, error) {
	return u.todoRes.GetTodos()
}

func (u *Usecase) GetTodoByID(id int64) (entityTodo.Todo, error) {
	return u.todoRes.GetTodoByID(id)
}

func (u *Usecase) AddTodo(todo entityTodo.Todo) error {
	if todo.Status == 0 {
		todo.Status = entityTodo.StatusActive
	}
	return u.todoRes.AddTodo(todo)
}

func (u *Usecase) DeleteTodo(id int64) error {
	return u.todoRes.DeleteTodo(id)
}

func (u *Usecase) UpdateTodo(req entityTodo.UpdateRequest, id int64) (int64, error) {
	return u.todoRes.UpdateTodo(id, req)
}
