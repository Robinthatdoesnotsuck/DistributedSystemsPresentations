package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"cityletterbox.com/pkg/discovery/consul"
	discovery "cityletterbox.com/pkg/registry"
	"cityletterbox.com/rating/internal/controller/rating"
	grpc_handler "cityletterbox.com/rating/internal/handler/grpc"
	"cityletterbox.com/rating/internal/repository/memory"
	"cityletterbox.com/src/gen"
	"google.golang.org/grpc"
)

const serviceName = "rating"

func main() {
	var port int
	flag.IntVar(&port, "port", 8082, "API handler port")
	flag.Parse()
	log.Printf("Starting rating service on port %d", port)
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
	repo := memory.New()
	ctrl := rating.New(repo)
	hdlr := grpc_handler.New(ctrl)
	list, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
	srv := grpc.NewServer()
	gen.RegisterRatingServiceServer(srv, hdlr)
	srv.Serve(list)
	if err := srv.Serve(list); err != nil {
		panic(err)
	}
}
