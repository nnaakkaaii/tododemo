package db

import (
	"context"
	"github.com/nnaakkaaii/tododemo/internal/todo"
)

type DB interface {
	PutTODO(ctx context.Context, t *todo.TODO) error
}
