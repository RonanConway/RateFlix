package main

import (
	"log"
	"net/http"

	"github.com/RonanConway/RateFlix/metadata/internal/controller/metadata"
	httphandler "github.com/RonanConway/RateFlix/metadata/internal/handler"
	memory "github.com/RonanConway/RateFlix/metadata/internal/repository/memory"
)

func main() {
	log.Println("Starting RateFlix")

	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)
	// Initialising all structures of the service then starting the http handler
	http.Handle("/metadata", http.HandlerFunc(h.GetMovieMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}

}
