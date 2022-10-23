package db

import (
	"context"
	"github.com/nnaakkaaii/tododemo/internal/model"
	"sync"
)

type memoryDB struct {
	db   map[string]*model.TODO
	lock sync.RWMutex
}

func NewMemoryDB() DB {
	return &memoryDB{db: map[string]*model.TODO{}}
}

func (m *memoryDB) SelectAllTODOs(ctx context.Context) ([]*model.TODO, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	result := make([]*model.TODO, 0, len(m.db))
	for _, t := range m.db {
		result = append(result, t)
	}

	return result, nil
}

func (m *memoryDB) InsertTODO(ctx context.Context, t *model.TODO) error {
	m.lock.Lock()
	m.db[t.ID] = t
	m.lock.Unlock()
	return nil
}
