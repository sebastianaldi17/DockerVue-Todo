package todo

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	entityTodo "github.com/sebastianaldi17/dockervue-todo/backend-go/internal/entity/todo"
)

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("but it works on local :)"))
}

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoUsecase.GetTodos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Printf("Error writing to HTTP: %+v", err)
	}
}

func (h *Handler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	get_id := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(get_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID must be integer"))
		return
	}

	todos, err := h.todoUsecase.GetTodoByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Printf("Error writing to HTTP: %+v", err)
	}
}

func (h *Handler) AddTodo(w http.ResponseWriter, r *http.Request) {
	var data entityTodo.Todo
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if len(data.Content) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("content must not be empty"))
		return
	}

	err := h.todoUsecase.AddTodo(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true}`))
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	delete_id := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(delete_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID must be integer"))
		return
	}

	if id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID must be bigger than 0"))
		return
	}

	err = h.todoUsecase.DeleteTodo(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true}`))
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var req entityTodo.UpdateRequest
	update_id := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(update_id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID must be integer"))
		return
	}

	if id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID must be bigger than 0"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	rows, err := h.todoUsecase.UpdateTodo(req, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if rows <= 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("ID not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true}`))
}
