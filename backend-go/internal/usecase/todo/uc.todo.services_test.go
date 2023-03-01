package todo

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	entityTodo "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/entity/todo"
)

var (
	errRes = errors.New("some error")
)

func TestUsecase_GetTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRes := NewMocktodoResource(ctrl)

	tests := []struct {
		name    string
		mockFn  func()
		want    []entityTodo.Todo
		wantErr bool
	}{
		{
			name: "success",
			mockFn: func() {
				mockTodoRes.
					EXPECT().
					GetTodos().
					Return([]entityTodo.Todo{{ID: 1}, {ID: 2}}, nil)
			},
			want:    []entityTodo.Todo{{ID: 1}, {ID: 2}},
			wantErr: false,
		},
		{
			name: "got error from resource",
			mockFn: func() {
				mockTodoRes.
					EXPECT().
					GetTodos().
					Return([]entityTodo.Todo{}, errRes)
			},
			want:    []entityTodo.Todo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				todoRes: mockTodoRes,
			}

			if tt.mockFn != nil {
				tt.mockFn()
			}

			got, err := u.GetTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_GetTodoByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRes := NewMocktodoResource(ctrl)

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
				id: 100,
			},
			mockFn: func(a args) {
				mockTodoRes.
					EXPECT().
					GetTodoByID(a.id).
					Return(entityTodo.Todo{ID: 100}, nil)
			},
			want:    entityTodo.Todo{ID: 100},
			wantErr: false,
		},
		{
			name: "got error from resource",
			args: args{
				id: 100,
			},
			mockFn: func(a args) {
				mockTodoRes.
					EXPECT().
					GetTodoByID(a.id).
					Return(entityTodo.Todo{}, errRes)
			},
			want:    entityTodo.Todo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				todoRes: mockTodoRes,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			got, err := u.GetTodoByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetTodoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetTodoByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_AddTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRes := NewMocktodoResource(ctrl)

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
					Content: "my first todo",
				},
			},
			mockFn: func(a args) {
				mockTodoRes.
					EXPECT().
					AddTodo(entityTodo.Todo{
						Content: "my first todo",
						Status:  entityTodo.StatusActive,
					})
			},
			wantErr: false,
		},
		{
			name: "got error from resource",
			args: args{
				todo: entityTodo.Todo{
					Content: "my first todo",
				},
			},
			mockFn: func(a args) {
				mockTodoRes.
					EXPECT().
					AddTodo(entityTodo.Todo{
						Content: "my first todo",
						Status:  entityTodo.StatusActive,
					}).
					Return(errRes)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				todoRes: mockTodoRes,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			if err := u.AddTodo(tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("Usecase.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecase_DeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRes := NewMocktodoResource(ctrl)

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
				id: 2,
			},
			mockFn: func(a args) {
				mockTodoRes.
					EXPECT().
					DeleteTodo(a.id)
			},
			wantErr: false,
		},
		{
			name: "got error from resource",
			args: args{
				id: 2,
			},
			mockFn: func(a args) {
				mockTodoRes.
					EXPECT().
					DeleteTodo(a.id).
					Return(errRes)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				todoRes: mockTodoRes,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			if err := u.DeleteTodo(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Usecase.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecase_UpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRes := NewMocktodoResource(ctrl)
	contentStr := "a"
	finishedBool := true
	statusInt := int32(entityTodo.StatusHidden)

	type args struct {
		req entityTodo.UpdateRequest
		id  int64
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
				id: 5,
				req: entityTodo.UpdateRequest{
					Content:  &contentStr,
					Finished: &finishedBool,
					Status:   &statusInt,
				},
			},
			mockFn: func(a args) {
				mockTodoRes.
					EXPECT().
					UpdateTodo(a.id, a.req).
					Return(int64(1), nil)
			},
			want:    int64(1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				todoRes: mockTodoRes,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			got, err := u.UpdateTodo(tt.args.req, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Usecase.UpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
