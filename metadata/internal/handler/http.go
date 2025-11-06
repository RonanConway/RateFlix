package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/RonanConway/RateFlix/metadata/internal/controller/metadata"
	"github.com/RonanConway/RateFlix/metadata/internal/repository"
	"github.com/gin-gonic/gin"
)

// Handler defines a movie metadata HTTP handler
type Handler struct {
	ctrl *metadata.Controller
}

// New creates a new movie metadata HTTP handler
func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) GetMovieMetadata(context *gin.Context) {
	Id, err := context.Param("id")

	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message", "No movie Id specifed in the request"})
	}

	metadata, err := h.ctrl.GetMetadata(context, Id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		context.JSON(http.StatusNotFound, gin.H{"message": "no metadata found for Id"})
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching from repository"})
		return
	}

	if err := encoder.Encode(metadata); err != nil {
		log.Printf("Response encode error")
	}

	return context.JSON(http.StatusOK, gin.H{"message": "Movie metadata ok"})
}
