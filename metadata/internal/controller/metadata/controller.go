package metadata

import (
	"context"
	"errors"

	"github.com/RonanConway/RateFlix/metadata/internal/repository"
	model "github.com/RonanConway/RateFlix/metadata/pkg"
)

var ErrNotFound = errors.New("not found")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller defines a metadata service controller
type Controller struct {
	repo metadataRepository
}

// New creates a metadata service controller
func New(repo metadataRepository) *Controller {
	return &Controller{repo}
}

func (controller *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	result, err := controller.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return result, nil
}
