package todo

import "net/http"

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hot reload is not working on Docker :)"))
}
