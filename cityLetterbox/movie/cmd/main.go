package cmd

import (
	"log"
	"net/http"

	"cityletterbox.com/movie/internal/controller/movie"
	metadatagateway "cityletterbox.com/movie/internal/gateway/metadata/http"
	ratinggateway "cityletterbox.com/movie/internal/gateway/rating/http"
	httphandler "cityletterbox.com/movie/internal/handler/http"
)

func main() {
	log.Println("Starting movie + rating service")
	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
