package handler

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/nnaakkaaii/tododemo/internal/model"
	"github.com/nnaakkaaii/tododemo/testing/db"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodosHandler_ServeHTTP(t *testing.T) {
	t.Parallel()

	todos := []*model.TODO{
		{
			ID:    "cfd8d91c-035d-4baf-9fef-41be5a9f4d61",
			Title: "write an esssay",
		},
	}
	jsonify := func(t interface{}) *bytes.Buffer {
		r := &bytes.Buffer{}
		if err := json.NewEncoder(r).Encode(t); err != nil {
			panic(err)
		}
		return r
	}

	tests := []struct {
		name     string
		reader   *http.Request
		wantCode int
		wantBody []byte
	}{
		{
			name:     "successful case: get",
			reader:   httptest.NewRequest(http.MethodGet, "/todos", nil),
			wantCode: http.StatusOK,
			wantBody: jsonify(todos).Bytes(),
		},
		{
			name:     "successful case: post",
			reader:   httptest.NewRequest(http.MethodPost, "/todos", jsonify(todos[0])),
			wantCode: http.StatusOK,
			wantBody: jsonify(todos[0]).Bytes(),
		},
	}

	d := db.NewFakeDB(todos)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			h := &todosHandler{db: d}
			rec := httptest.NewRecorder()

			h.ServeHTTP(rec, tc.reader)

			if diff := cmp.Diff(tc.wantCode, rec.Code); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(tc.wantBody, rec.Body.Bytes()); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
