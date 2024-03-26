package transport

import (
	"context"
	pb "hashing/pkg/hasher"
	"hashing/service"
)

func NewGrpcServer(svc *service.HashingService) pb.HashingServiceServer {
	return &grpcServer{service: svc}
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
