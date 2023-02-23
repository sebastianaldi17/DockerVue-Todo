package todo

import (
	"fmt"

	entityTodo "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/entity/todo"
)

func (r *Resource) GetTodos() ([]entityTodo.Todo, error) {
	todos := make([]entityTodo.Todo, 0)
	err := r.db.Select(&todos, queryGetAll)

	return todos, err
}

func (r *Resource) AddTodo(todo entityTodo.Todo) error {
	_, err := r.db.Exec(queryAddTodo, todo.Content, todo.Status, todo.Finished)
	return err
}

func (r *Resource) DeleteTodo(id int64) error {
	_, err := r.db.Exec(queryDeleteTodo, id)
	return err
}

func (r *Resource) UpdateTodo(id int64, req entityTodo.UpdateRequest) (int64, error) {
	updateVals := []interface{}{}
	updateFields := []string{}

	if req.Content != nil {
		updateVals = append(updateVals, *req.Content)
		updateFields = append(updateFields, "content")
	}

	if req.Finished != nil {
		updateVals = append(updateVals, *req.Finished)
		updateFields = append(updateFields, "finished")
	}

	if req.Status != nil {
		updateVals = append(updateVals, *req.Status)
		updateFields = append(updateFields, "status")
	}

	sets := ""
	for _, field := range updateFields {
		sets += fmt.Sprintf("%s = ?,", field)
	}

	updateVals = append(updateVals, id)

	query := fmt.Sprintf(queryUpdateTodo, sets)
	query = r.db.Rebind(query)
	res, err := r.db.Exec(query, updateVals...)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected, err
}
