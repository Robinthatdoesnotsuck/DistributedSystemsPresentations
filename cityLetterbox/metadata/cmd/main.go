package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"cityletterbox.com/metadata/internal/controller/metadata"
	grpc_handler "cityletterbox.com/metadata/internal/handler/grpc"
	"cityletterbox.com/metadata/internal/repository/memory"
	"cityletterbox.com/pkg/discovery/consul"
	discovery "cityletterbox.com/pkg/registry"
	"cityletterbox.com/src/gen"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

const serviceName = "metadata"

func main() {
	f, err := os.Open("base.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg serviceConfig
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}
	port := cfg.APIConfig.Port
	log.Printf("Starting metadata service on port: %d", port)
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
	ctrl := metadata.New(repo)
	hdlr := grpc_handler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.APIConfig.Port))
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
	srv := grpc.NewServer()
	gen.RegisterMetadataServiceServer(srv, hdlr)
	srv.Serve(lis)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
