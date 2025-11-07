package httphandler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/RonanConway/RateFlix/metadata/internal/controller/metadata"
	"github.com/RonanConway/RateFlix/metadata/internal/repository"
)

// Handler defines a movie metadata HTTP handler
type Handler struct {
	ctrl *metadata.Controller
}

// New creates a new movie metadata HTTP handler
func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) GetMovieMetadata(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	m, err := h.ctrl.GetMetadata(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("REpository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}

}
