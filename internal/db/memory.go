package db

import (
	"context"
	"github.com/nnaakkaaii/tododemo/internal/todo"
	"sync"
)

var _ DB = (*memoryDB)(nil)

type memoryDB struct {
	db   map[string]*todo.TODO
	lock sync.RWMutex
}

func (m *memoryDB) PutTODO(ctx context.Context, t *todo.TODO) error {
	m.lock.Lock()
	m.db[t.ID] = t
	m.lock.Unlock()
	return nil
}
