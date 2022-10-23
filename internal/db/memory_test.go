package db

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/nnaakkaaii/tododemo/internal/model"
	"testing"
)

func TestMemoryDB_SelectAllTODOs(t *testing.T) {
	t.Parallel()

	todos := []*model.TODO{
		{
			ID:    "d3a6633e-0409-4fba-98be-de918e0fbbcc",
			Title: "apply for kaggle",
		},
	}

	tests := []struct {
		name       string
		db         map[string]*model.TODO
		wantEntity []*model.TODO
		wantErr    bool
	}{
		{
			name:       "successful case",
			db:         map[string]*model.TODO{todos[0].ID: todos[0]},
			wantEntity: todos,
			wantErr:    false,
		},
	}

	ctx := context.Background()

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			d := &memoryDB{db: tc.db}

			entity, err := d.SelectAllTODOs(ctx)

			if diff := cmp.Diff(tc.wantErr, err != nil); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(tc.wantEntity, entity); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestMemoryDB_InsertTODO(t *testing.T) {
	t.Parallel()

	todos := []*model.TODO{
		{
			ID:    "3bfef2d7-3f3f-4f9a-81d9-28bb61f33ef0",
			Title: "mail to the professor",
		},
	}

	tests := []struct {
		name    string
		entity  *model.TODO
		db      map[string]*model.TODO
		wantDB  map[string]*model.TODO
		wantErr bool
	}{
		{
			name:    "successful case",
			entity:  todos[0],
			db:      map[string]*model.TODO{},
			wantDB:  map[string]*model.TODO{todos[0].ID: todos[0]},
			wantErr: false,
		},
	}

	ctx := context.Background()

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			d := &memoryDB{db: tc.db}

			err := d.InsertTODO(ctx, tc.entity)

			if diff := cmp.Diff(tc.wantErr, err != nil); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(tc.wantDB, d.db); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
