package todo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	entityTodo "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/entity/todo"
)

func TestHandler_Hello(t *testing.T) {
	rr := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantBody string
		wantCode int
	}{
		{
			name: "success",
			args: args{
				w: rr,
				r: &http.Request{
					Method: http.MethodGet,
				},
			},
			wantBody: "Hi",
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{}

			h.Hello(tt.args.w, tt.args.r)

			result := rr.Result()
			if result.StatusCode != tt.wantCode {
				t.Errorf("Handler.Hello() different http code expected = %v, got %v", tt.wantCode, result.StatusCode)
			}
			responseByte, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Errorf("Handler.Hello() return error on read: %v", err)
			}
			if string(responseByte) != tt.wantBody {
				t.Errorf("Handler.Hello() different body expected %v, got %v", tt.wantBody, string(responseByte))
			}
		})
	}
}

func TestHandler_GetTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoUC := NewMocktodoUsecase(ctrl)

	returnedList := []entityTodo.Todo{{ID: 1}, {ID: 2}}
	marshalledString, _ := json.Marshal(returnedList)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		mockFn   func(a args)
		wantBody string
		wantCode int
	}{
		{
			name: "success",
			args: args{
				r: &http.Request{
					Method: http.MethodGet,
				},
			},
			mockFn: func(a args) {
				mockTodoUC.
					EXPECT().
					GetTodos().
					Return(returnedList, nil)
			},
			wantBody: string(marshalledString),
			wantCode: http.StatusOK,
		},
		{
			name: "got error from usecase",
			args: args{
				r: &http.Request{
					Method: http.MethodGet,
				},
			},
			mockFn: func(a args) {
				mockTodoUC.
					EXPECT().
					GetTodos().
					Return(nil, errors.New("some error"))
			},
			wantBody: "some error",
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			h := &Handler{
				todoUsecase: mockTodoUC,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			h.GetTodos(rr, tt.args.r)

			result := rr.Result()
			if result.StatusCode != tt.wantCode {
				t.Errorf("Handler.GetTodos() different http code expected = %v, got %v", tt.wantCode, result.StatusCode)
			}
			responseByte, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Errorf("Handler.GetTodos() return error on read: %v", err)
			}
			if string(responseByte) != tt.wantBody {
				t.Errorf("Handler.GetTodos() different body expected %v, got %v", tt.wantBody, string(responseByte))
			}
		})
	}
}

func TestHandler_DeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoUC := NewMocktodoUsecase(ctrl)

	req := httptest.NewRequest(http.MethodDelete, "/todo/{id}", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	badreq := httptest.NewRequest(http.MethodDelete, "/todo/{id}", nil)
	badrctx := chi.NewRouteContext()
	badrctx.URLParams.Add("id", "-999")
	badreq = badreq.WithContext(context.WithValue(badreq.Context(), chi.RouteCtxKey, badrctx))

	badreq2 := httptest.NewRequest(http.MethodDelete, "/todo/{id}", nil)
	badrctx2 := chi.NewRouteContext()
	badrctx2.URLParams.Add("id", "string")
	badreq2 = badreq2.WithContext(context.WithValue(badreq2.Context(), chi.RouteCtxKey, badrctx2))

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		mockFn   func(a args)
		wantBody string
		wantCode int
	}{
		{
			name: "success",
			args: args{
				r: req,
			},
			mockFn: func(a args) {
				mockTodoUC.
					EXPECT().
					DeleteTodo(int64(123)).
					Return(nil)
			},
			wantBody: `{"success": true}`,
			wantCode: http.StatusOK,
		},
		{
			name: "error from usecase",
			args: args{
				r: req,
			},
			mockFn: func(a args) {
				mockTodoUC.
					EXPECT().
					DeleteTodo(int64(123)).
					Return(errors.New("some error"))
			},
			wantBody: "some error",
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "bad request (negative number)",
			args: args{
				r: badreq,
			},
			wantBody: "ID must be bigger than 0",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "bad request (string)",
			args: args{
				r: badreq2,
			},
			wantBody: "ID must be integer",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			h := &Handler{
				todoUsecase: mockTodoUC,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			h.DeleteTodo(rr, tt.args.r)

			result := rr.Result()
			if result.StatusCode != tt.wantCode {
				t.Errorf("Handler.DeleteTodo() different http code expected = %v, got %v", tt.wantCode, result.StatusCode)
			}
			responseByte, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Errorf("Handler.DeleteTodo() return error on read: %v", err)
			}
			if string(responseByte) != tt.wantBody {
				t.Errorf("Handler.DeleteTodo() different body expected %v, got %v", tt.wantBody, string(responseByte))
			}
		})
	}
}

func TestHandler_GetTodoByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoUC := NewMocktodoUsecase(ctrl)

	req := httptest.NewRequest(http.MethodGet, "/todo/{id}", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	badreq := httptest.NewRequest(http.MethodGet, "/todo/{id}", nil)
	badrctx := chi.NewRouteContext()
	badrctx.URLParams.Add("id", "-999")
	badreq = badreq.WithContext(context.WithValue(badreq.Context(), chi.RouteCtxKey, badrctx))

	badreq2 := httptest.NewRequest(http.MethodGet, "/todo/{id}", nil)
	badrctx2 := chi.NewRouteContext()
	badrctx2.URLParams.Add("id", "string")
	badreq2 = badreq.WithContext(context.WithValue(badreq2.Context(), chi.RouteCtxKey, badrctx2))

	todo := entityTodo.Todo{ID: 123}
	marshalled, _ := json.Marshal(todo)

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		mockFn   func(a args)
		wantBody string
		wantCode int
	}{
		{
			name: "success",
			args: args{
				r: req,
			},
			mockFn: func(a args) {
				mockTodoUC.
					EXPECT().
					GetTodoByID(int64(123)).
					Return(todo, nil)
			},
			wantBody: string(marshalled),
			wantCode: http.StatusOK,
		},
		{
			name: "error on usecase",
			args: args{
				r: req,
			},
			mockFn: func(a args) {
				mockTodoUC.
					EXPECT().
					GetTodoByID(int64(123)).
					Return(todo, errors.New("some error"))
			},
			wantBody: "some error",
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "bad request (negative number)",
			args: args{
				r: badreq,
			},
			wantBody: "ID must be bigger than 0",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "bad request (string)",
			args: args{
				r: badreq2,
			},
			wantBody: "ID must be integer",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			h := &Handler{
				todoUsecase: mockTodoUC,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			h.GetTodoByID(rr, tt.args.r)

			result := rr.Result()
			if result.StatusCode != tt.wantCode {
				t.Errorf("Handler.GetTodoByID() different http code expected = %v, got %v", tt.wantCode, result.StatusCode)
			}
			responseByte, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Errorf("Handler.GetTodoByID() return error on read: %v", err)
			}
			if string(responseByte) != tt.wantBody {
				t.Errorf("Handler.GetTodoByID() different body expected %v, got %v", tt.wantBody, string(responseByte))
			}
		})
	}
}

func TestHandler_AddTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoUC := NewMocktodoUsecase(ctrl)

	todoEnt := entityTodo.Todo{Content: "new todo"}
	todoBody, _ := json.Marshal(todoEnt)

	emptyEnt := entityTodo.Todo{}
	emptyBody, _ := json.Marshal(emptyEnt)

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		mockFn   func(a args)
		wantBody string
		wantCode int
	}{
		{
			name: "success",
			args: args{
				r: httptest.NewRequest(http.MethodPost, "/todo", bytes.NewReader(todoBody)),
			},
			mockFn: func(a args) {
				mockTodoUC.
					EXPECT().
					AddTodo(todoEnt).
					Return(nil)
			},
			wantBody: `{"success": true}`,
			wantCode: http.StatusOK,
		},
		{
			name: "error on usecase",
			args: args{
				r: httptest.NewRequest(http.MethodPost, "/todo", bytes.NewReader(todoBody)),
			},
			mockFn: func(a args) {
				mockTodoUC.
					EXPECT().
					AddTodo(todoEnt).
					Return(errors.New("some error"))
			},
			wantBody: "some error",
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "bad body request",
			args: args{
				r: httptest.NewRequest(http.MethodPost, "/todo", bytes.NewReader([]byte("asdf"))),
			},
			wantBody: "invalid character 'a' looking for beginning of value",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "empty content",
			args: args{
				r: httptest.NewRequest(http.MethodPost, "/todo", bytes.NewReader(emptyBody)),
			},
			wantBody: "content must not be empty",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			h := &Handler{
				todoUsecase: mockTodoUC,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			h.AddTodo(rr, tt.args.r)

			result := rr.Result()
			if result.StatusCode != tt.wantCode {
				t.Errorf("Handler.AddTodo() different http code expected = %v, got %v", tt.wantCode, result.StatusCode)
			}
			responseByte, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Errorf("Handler.AddTodo() return error on read: %v", err)
			}
			if string(responseByte) != tt.wantBody {
				t.Errorf("Handler.AddTodo() different body expected %v, got %v", tt.wantBody, string(responseByte))
			}
		})
	}
}

func TestHandler_UpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoUC := NewMocktodoUsecase(ctrl)

	updateContent := "someContent"
	updateStatus := int32(entityTodo.StatusHidden)
	updateFinished := true
	updateReq := entityTodo.UpdateRequest{
		Content:  &updateContent,
		Status:   &updateStatus,
		Finished: &updateFinished,
	}
	updateBody, _ := json.Marshal(updateReq)

	req := httptest.NewRequest(http.MethodPut, "/todo/{id}", bytes.NewReader(updateBody))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	req2 := httptest.NewRequest(http.MethodPut, "/todo/{id}", bytes.NewReader(updateBody))
	rctx2 := chi.NewRouteContext()
	rctx2.URLParams.Add("id", "123")
	req2 = req2.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx2))

	badreq := httptest.NewRequest(http.MethodPut, "/todo/{id}", nil)
	badrctx := chi.NewRouteContext()
	badrctx.URLParams.Add("id", "-999")
	badreq = badreq.WithContext(context.WithValue(badreq.Context(), chi.RouteCtxKey, badrctx))

	badreq2 := httptest.NewRequest(http.MethodPut, "/todo/{id}", nil)
	badrctx2 := chi.NewRouteContext()
	badrctx2.URLParams.Add("id", "string")
	badreq2 = badreq2.WithContext(context.WithValue(badreq2.Context(), chi.RouteCtxKey, badrctx2))

	badbodyreq := httptest.NewRequest(http.MethodPut, "/todo/{id}", bytes.NewReader([]byte("asdf")))
	badbodyctx := chi.NewRouteContext()
	badbodyctx.URLParams.Add("id", "123")
	badbodyreq = badbodyreq.WithContext(context.WithValue(badbodyreq.Context(), chi.RouteCtxKey, badbodyctx))

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		mockFn   func(a args)
		wantBody string
		wantCode int
	}{
		{
			name: "success",
			args: args{
				r: req,
			},
			mockFn: func(a args) {
				mockTodoUC.
					EXPECT().
					UpdateTodo(updateReq, int64(123)).
					Return(int64(1), nil)
			},
			wantBody: `{"success": true}`,
			wantCode: http.StatusOK,
		},
		{
			name: "bad request (negative number)",
			args: args{
				r: badreq,
			},
			wantBody: `ID must be bigger than 0`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "bad request (string)",
			args: args{
				r: badreq2,
			},
			wantBody: `ID must be integer`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "bad body request",
			args: args{
				r: badbodyreq,
			},
			wantBody: `invalid character 'a' looking for beginning of value`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "error on usecase",
			args: args{
				r: req2,
			},
			mockFn: func(a args) {
				mockTodoUC.
					EXPECT().
					UpdateTodo(updateReq, int64(123)).
					Return(int64(0), errors.New("some error"))
			},
			wantBody: `some error`,
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			h := &Handler{
				todoUsecase: mockTodoUC,
			}

			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			h.UpdateTodo(rr, tt.args.r)

			result := rr.Result()
			if result.StatusCode != tt.wantCode {
				t.Errorf("Handler.UpdateTodo() different http code expected = %v, got %v", tt.wantCode, result.StatusCode)
			}
			responseByte, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Errorf("Handler.UpdateTodo() return error on read: %v", err)
			}
			if string(responseByte) != tt.wantBody {
				t.Errorf("Handler.UpdateTodo() different body expected %v, got %v", tt.wantBody, string(responseByte))
			}
		})
	}
}
