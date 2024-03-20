package main

import (
	"final_project/internal/hashing/service"
	"final_project/internal/hashing/transport"
	pb "final_project/pkg/hasher"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := "50052"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hashingService := service.NewHashingService()
	pb.RegisterHashingServiceServer(s, transport.NewGrpcServer(hashingService))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
