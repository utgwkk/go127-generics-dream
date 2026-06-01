package apiserver

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	mux *http.ServeMux
}

func NewHandler() *Handler {
	return &Handler{
		mux: http.NewServeMux(),
	}
}

func (h *Handler) Handle[In any](pattern string, f func(w http.ResponseWriter, r *http.Request, data In)) *Handler {
	h.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		var data In
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		f(w, r, data)
	})
	return h
}
