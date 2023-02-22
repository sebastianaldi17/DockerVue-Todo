package todo

import "database/sql"

type Resource struct {
	db *sql.DB
}

func New(db *sql.DB) *Resource {
	todo := Resource{
		db: db,
	}
	return &todo
}
