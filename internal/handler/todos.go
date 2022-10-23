package handler

import (
	"encoding/json"
	"github.com/nnaakkaaii/tododemo/internal/db"
	"github.com/nnaakkaaii/tododemo/internal/model"
	"net/http"
)

type todosHandler struct {
	db db.DB
}

func NewTODOsHandler(db db.DB) http.Handler {
	return &todosHandler{db: db}
}

func (h *todosHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		todos, err := h.db.SelectAllTODOs(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&todos); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var t model.TODO
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if t.ID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.db.InsertTODO(r.Context(), &t); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&t); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
