package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"cityletterbox.com/movie/internal/controller/movie"
	metadatagateway "cityletterbox.com/movie/internal/gateway/metadata/http"
	ratinggateway "cityletterbox.com/movie/internal/gateway/rating/http"
	httphandler "cityletterbox.com/movie/internal/handler/http"
	"cityletterbox.com/pkg/discovery/consul"
	discovery "cityletterbox.com/pkg/registry"
)

const serviceName = "movie"

func main() {
	var port int
	flag.IntVar(&port, "port", 8083, "API Handler port")
	flag.Parse()
	log.Printf("Starting movie + rating service on port: %d", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
