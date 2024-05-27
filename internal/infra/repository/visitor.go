package infra_repository

import (
	"sync"

	"github.com/salvovitale/dddeu24-tact-patterns-ws/internal/domain"
)

type VisitorInMemoryRepository struct {
	visitors map[string]domain.Visitor
	mu       sync.Mutex
}

func NewVisitorInMemoryRepository() *VisitorInMemoryRepository {
	return &VisitorInMemoryRepository{
		visitors: make(map[string]domain.Visitor),
		mu:       sync.Mutex{},
	}
}

func (r *VisitorInMemoryRepository) Get(id string) (domain.Visitor, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	v, ok := r.visitors[id]
	if !ok {
		return domain.Visitor{}, domain.ErrVisitorNotFound
	}
	return v, nil
}

func (r *VisitorInMemoryRepository) Save(v domain.Visitor) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.visitors[v.ID] = v
	return nil
}

func (r *VisitorInMemoryRepository) Clear() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.visitors = make(map[string]domain.Visitor)
	return nil
}
