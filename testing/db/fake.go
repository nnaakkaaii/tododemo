package db

import (
	"context"
	"github.com/nnaakkaaii/tododemo/internal/db"
	"github.com/nnaakkaaii/tododemo/internal/model"
)

type fakeDB struct {
	seedTODOs []*model.TODO
}

func NewFakeDB(t []*model.TODO) db.DB {
	return &fakeDB{seedTODOs: t}
}

func (f *fakeDB) SelectAllTODOs(ctx context.Context) ([]*model.TODO, error) {
	return f.seedTODOs, nil
}

func (f *fakeDB) InsertTODO(ctx context.Context, t *model.TODO) error {
	return nil
}
