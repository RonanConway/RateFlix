package main

import (
	"log"
	"net/http"

	rating "github.com/RonanConway/RateFlix/rating/internal/controller"
	httphandler "github.com/RonanConway/RateFlix/rating/internal/handler"
	"github.com/RonanConway/RateFlix/rating/internal/repository/memory"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}

}
