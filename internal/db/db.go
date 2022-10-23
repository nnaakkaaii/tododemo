package db

import (
	"context"
	"github.com/nnaakkaaii/tododemo/internal/model"
)

type DB interface {
	SelectAllTODOs(ctx context.Context) ([]*model.TODO, error)
	InsertTODO(ctx context.Context, t *model.TODO) error
}
