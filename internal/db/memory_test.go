package db

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/nnaakkaaii/tododemo/internal/todo"
)

func TestMemoryDB_PutTODO(t *testing.T) {
	t.Parallel()

	todo1 := &todo.TODO{
		ID:    "43b9c5c8-bb77-416e-8f65-22a23d221d64",
		Title: "brush the gopher",
	}
	tests := map[string]struct {
		todo     *todo.TODO
		expected map[string]*todo.TODO
	}{
		"put": {
			todo:     todo1,
			expected: map[string]*todo.TODO{todo1.ID: todo1},
		},
	}

	ctx := context.Background()
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			d := &memoryDB{db: map[string]*todo.TODO{}}
			if err := d.PutTODO(ctx, test.todo); err != nil {
				t.Fatalf("failed to put a todo: %s", err.Error())
			}

			if diff := cmp.Diff(test.expected, d.db); diff != "" {
				t.Errorf("\n(-expected, +actual)\n%s", diff)
			}
		})
	}
}
