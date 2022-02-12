package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kikils/backend/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9999))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := proto.Server{}
	grpcServer := grpc.NewServer()
	// serverにserviceを追加
	proto.RegisterAddNumServiceServer(grpcServer, &s)
	reflection.Register(grpcServer)


	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	} else {
		log.Printf("Server started!!")
	}
	log.Printf("here")
}