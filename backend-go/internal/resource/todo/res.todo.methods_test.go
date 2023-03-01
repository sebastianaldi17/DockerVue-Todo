package todo

import (
	"reflect"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	entityTodo "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/entity/todo"
)

func TestResource_GetTodos(t *testing.T) {
	mockdb, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error when initializing sqlmock: %v", err)
	}
	sqlxdb := sqlx.NewDb(mockdb, "sqlmock")

	tests := []struct {
		name    string
		mockFn  func()
		want    []entityTodo.Todo
		wantErr bool
	}{
		{
			name: "success",
			mockFn: func() {
				mocksql.
					ExpectQuery(queryGetAll).
					WillReturnRows(sqlmock.NewRows([]string{"id", "content"}).
						AddRow(1, "hi").
						AddRow(2, "hello"))
			},
			want: []entityTodo.Todo{{
				ID:      1,
				Content: "hi",
			}, {
				ID:      2,
				Content: "hello",
			}},
		},
		{
			name: "db error",
			mockFn: func() {
				mocksql.
					ExpectQuery("SELECT").
					WillReturnError(sqlmock.ErrCancelled)
			},
			want:    []entityTodo.Todo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Resource{
				db: sqlxdb,
			}

			if tt.mockFn != nil {
				tt.mockFn()
			}

			got, err := r.GetTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("Resource.GetTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resource.GetTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResource_GetTodoByID(t *testing.T) {
	mockdb, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error when initializing sqlmock: %v", err)
	}
	sqlxdb := sqlx.NewDb(mockdb, "sqlmock")

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(a args)
		want    entityTodo.Todo
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			mockFn: func(a args) {
				mocksql.
					ExpectQuery("SELECT").
					WithArgs(a.id).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "content"}).
							AddRow(1, "sample text"))
			},
			want: entityTodo.Todo{
				ID:      1,
				Content: "sample text",
			},
		},
		{
			name: "db error",
			args: args{
				id: 1,
			},
			mockFn: func(a args) {
				mocksql.
					ExpectQuery("SELECT").
					WithArgs(a.id).
					WillReturnError(sqlmock.ErrCancelled)
			},
			want:    entityTodo.Todo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Resource{
				db: sqlxdb,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			got, err := r.GetTodoByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resource.GetTodoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resource.GetTodoByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResource_AddTodo(t *testing.T) {
	mockdb, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error when initializing sqlmock: %v", err)
	}
	sqlxdb := sqlx.NewDb(mockdb, "sqlmock")

	type args struct {
		todo entityTodo.Todo
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(a args)
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				todo: entityTodo.Todo{
					Content:  "content",
					Status:   entityTodo.StatusActive,
					Finished: false,
				},
			},
			mockFn: func(a args) {
				mocksql.
					ExpectExec("INSERT").
					WithArgs(a.todo.Content, a.todo.Status, a.todo.Finished).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "db error",
			args: args{
				todo: entityTodo.Todo{
					Content:  "content",
					Status:   entityTodo.StatusActive,
					Finished: false,
				},
			},
			mockFn: func(a args) {
				mocksql.
					ExpectExec("INSERT").
					WithArgs(a.todo.Content, a.todo.Status, a.todo.Finished).
					WillReturnError(sqlmock.ErrCancelled)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Resource{
				db: sqlxdb,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			if err := r.AddTodo(tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("Resource.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResource_DeleteTodo(t *testing.T) {
	mockdb, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error when initializing sqlmock: %v", err)
	}
	sqlxdb := sqlx.NewDb(mockdb, "sqlmock")

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(a args)
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			mockFn: func(a args) {
				mocksql.
					ExpectExec("DELETE").
					WithArgs(a.id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "db error",
			args: args{
				id: 1,
			},
			mockFn: func(a args) {
				mocksql.
					ExpectExec("DELETE").
					WithArgs(a.id).
					WillReturnError(sqlmock.ErrCancelled)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Resource{
				db: sqlxdb,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			if err := r.DeleteTodo(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Resource.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResource_UpdateTodo(t *testing.T) {
	mockdb, mocksql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error when initializing sqlmock: %v", err)
	}
	sqlxdb := sqlx.NewDb(mockdb, "sqlmock")

	contentStr := "some content"
	statusInt := int32(entityTodo.StatusHidden)
	finishedBool := true

	type args struct {
		id  int64
		req entityTodo.UpdateRequest
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(a args)
		want    int64
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
				req: entityTodo.UpdateRequest{
					Content:  &contentStr,
					Status:   &statusInt,
					Finished: &finishedBool,
				},
			},
			mockFn: func(a args) {
				mocksql.
					ExpectExec("UPDATE").
					WithArgs(contentStr, finishedBool, statusInt, a.id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: 1,
		},
		{
			name: "db error",
			args: args{
				id: 1,
				req: entityTodo.UpdateRequest{
					Content: &contentStr,
				},
			},
			mockFn: func(a args) {
				mocksql.
					ExpectExec("UPDATE").
					WithArgs(contentStr, a.id).
					WillReturnError(sqlmock.ErrCancelled)
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Resource{
				db: sqlxdb,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			got, err := r.UpdateTodo(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resource.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Resource.UpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
