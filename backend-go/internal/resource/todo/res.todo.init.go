package todo

import (
	"github.com/jmoiron/sqlx"
)

type Resource struct {
	db *sqlx.DB
}

const (
	queryGetAll = `
	SELECT
		id,
		content,
		status,
		finished,
		created_at,
		updated_at
	FROM
		todos
	`

	queryGetByID = `
	SELECT
		id,
		content,
		status,
		finished,
		created_at,
		updated_at
	FROM
		todos
	WHERE
		id = $1
	`

	queryAddTodo = `
	INSERT INTO
		todos(content, status, finished)
	VALUES
		($1, $2, $3)
	`

	queryDeleteTodo = `
	DELETE FROM
		todos
	WHERE
		id = $1
	`

	queryUpdateTodo = `
	UPDATE
		todos
	SET
		%s
		updated_at = now()
	WHERE
		id = ?
	`
)

func New(db *sqlx.DB) *Resource {
	todo := Resource{
		db: db,
	}
	return &todo
}
