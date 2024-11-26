package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "paper-news-backend/gen/pb/api"
	c "paper-news-backend/internal/controller"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// make a listener
	serverPort := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		panic(err)
	}

	// make a grpc server
	s := grpc.NewServer()

	// register services
	pb.RegisterGreetingServiceServer(s, c.NewGreetingServiceServer())

	reflection.Register(s)

	go func() {
		// start grpc server
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}()

	// Wait for a termination signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down gRPC server...")
	s.GracefulStop()
}
