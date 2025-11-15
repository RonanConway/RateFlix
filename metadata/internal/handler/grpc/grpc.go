package grpc

import (
	"context"
	"errors"

	"github.com/RonanConway/RateFlix/gen"
	"github.com/RonanConway/RateFlix/metadata/internal/controller/metadata"
	"github.com/RonanConway/RateFlix/metadata/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler defines a movie metadata gRPC handler.
type Handler struct {
	gen.UnimplementedMetadataServiceServer
	ctrl *metadata.Controller
}

// New creates a new movie metadata gRPC handler.
func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetMetadataByID returns movie metadata by id.
func (h *Handler) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	if req == nil || req.MovieId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.Get(ctx, req.MovieId)
	if err != nil && errors.Is(err, metadata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// Using the MetadataToProto mapping function to transform
	// the internal structures into the generated ones.
	return &gen.GetMetadataResponse{Metadata: model.MetadataToProto(m)}, nil
}
