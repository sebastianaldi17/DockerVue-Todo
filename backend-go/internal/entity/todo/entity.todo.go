package todo

import "time"

type Todo struct {
	ID        int64     `json:"id" db:"id"`
	Content   string    `json:"content" db:"content"`
	Status    int32     `json:"status" db:"status"`
	Finished  bool      `json:"finished" db:"finished"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UpdateRequest struct {
	Content  *string `json:"content"`
	Status   *int32  `json:"status"`
	Finished *bool   `json:"finished"`
}
