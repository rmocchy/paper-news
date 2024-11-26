package controller

import (
	"context"
	pb "paper-news-backend/gen/pb/api"
)

func NewGreetingServiceServer() *greetingServiceServer {
	return &greetingServiceServer{}
}

type greetingServiceServer struct {
	pb.UnimplementedGreetingServiceServer
}

func (s *greetingServiceServer) Greet(ctx context.Context, req *pb.GreetRequestPayload) (*pb.GreetResponsePayload, error) {
	return &pb.GreetResponsePayload{
		Message: "Hello, " + req.Name,
	}, nil
}
