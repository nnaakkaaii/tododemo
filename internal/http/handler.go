package http

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nnaakkaaii/tododemo/internal/db"
	"github.com/nnaakkaaii/tododemo/internal/todo"
	"net/http"
)

var _ http.Handler = (*createHandler)(nil)

type createHandler struct {
	db db.DB
}

func (h *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var t todo.TODO
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t.ID = uuid.New().String()

	if err := h.db.PutTODO(r.Context(), &t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
