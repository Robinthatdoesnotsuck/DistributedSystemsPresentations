package main

import (
	"log"
	"net"

	"cityletterbox.com/metadata/internal/controller/metadata"
	grpc_handler "cityletterbox.com/metadata/internal/handler/grpc"
	"cityletterbox.com/metadata/internal/repository/memory"
	"cityletterbox.com/src/gen"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting metadata movie service with grpc")
	repo := memory.New()
	ctrl := metadata.New(repo)
	hdlr := grpc_handler.New(ctrl)
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
	srv := grpc.NewServer()
	gen.RegisterMetadataServiceServer(srv, hdlr)
	srv.Serve(lis)
}
