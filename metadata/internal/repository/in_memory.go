package in_memory

import (
	"context"
	"sync"

	"github.com/RonanConway/RateFlix/metadata/internal/repository"
	model "github.com/RonanConway/RateFlix/metadata/pkg"
)

// Repository defines a in memory movie metadata repository
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New creates a new in memory repo
func New() *Repository {
	return &Repository{data: map[string]*model.Metadata}
}

// Retrieves movie metadata for movie by Id
// Good practice for all functions performing IO operations to accept a context.
func (r *Repository) GetMetadata(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	metadata, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return metadata, nil
}

// Adds movie metadata for a given Id
func (r *Repository) AddMetadata(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
