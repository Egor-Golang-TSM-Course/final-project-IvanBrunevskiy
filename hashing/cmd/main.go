package main

import (
	"google.golang.org/grpc"
	pb "hashing/pkg/hasher"
	"hashing/service"
	"hashing/transport"
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
