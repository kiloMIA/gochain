package main

import (
	"context"
	pb "github.com/kiloMIA/cmd/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedNodeServiceServer
}

func (s *Server) ProposeContract(ctx context.Context, req *pb.ProposeContractRequest) (*pb.ProposeContractResponse, error) {

	return &pb.ProposeContractResponse{Accepted: true}, nil
}

func (s *Server) ValidateContract(ctx context.Context, req *pb.ValidateContractRequest) (*pb.ValidateContractResponse, error) {

	return &pb.ValidateContractResponse{Valid: true}, nil
}

func (s *Server) AchieveConsensus(ctx context.Context, req *pb.ConsensusRequest) (*pb.ConsensusResponse, error) {

	return &pb.ConsensusResponse{Agreed: true}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterNodeServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
