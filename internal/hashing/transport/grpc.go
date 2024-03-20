package transport

import (
	"context"
	"net"

	"final_project/internal/hashing/service"
	pb "final_project/pkg/hasher"
	"google.golang.org/grpc"
)

func NewGrpcServer(svc *service.HashingService) pb.HashingServiceServer {
	return &grpcServer{service: svc}
}

func StartGRPCServer(svc *service.HashingService, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterHashingServiceServer(server, &grpcServer{service: svc})
	return server.Serve(lis)
}

type grpcServer struct {
	pb.UnimplementedHashingServiceServer
	service *service.HashingService
}

func (s *grpcServer) CheckHash(ctx context.Context, req *pb.HashRequest) (*pb.HashResponse, error) {
	exists, err := s.service.CheckHash(req.Payload)
	if err != nil {
		return nil, err
	}
	return &pb.HashResponse{Exists: exists}, nil
}

func (s *grpcServer) GetHash(ctx context.Context, req *pb.HashRequest) (*pb.HashResponse, error) {
	hash, err := s.service.GetHash(req.Payload)
	if err != nil {
		return nil, err
	}
	return &pb.HashResponse{Hash: hash}, nil
}

func (s *grpcServer) CreateHash(ctx context.Context, req *pb.HashRequest) (*pb.HashResponse, error) {
	hash, err := s.service.CreateHash(req.Payload)
	if err != nil {
		return nil, err
	}
	return &pb.HashResponse{Hash: hash}, nil
}
