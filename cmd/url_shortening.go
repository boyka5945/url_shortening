package main

import (
	"context"
	"log"
	"net"
	"url_shortening/config"
	"url_shortening/infra"
	pb "url_shortening/model"
	"url_shortening/service"

	"google.golang.org/grpc"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = infra.InitLog()
	if err != nil {
		log.Fatalf("failed to initialize log: %v", err)
	}

	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(loggingInterceptor),
	}
	// Create a new gRPC server
	server := grpc.NewServer(serverOptions...)

	// Register the UserService server with the gRPC server
	pb.RegisterUrlShorteningServiceServer(server, service.NewUserServiceServer())
	// Start the gRPC server and listen for incoming requests
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// Log the incoming request
	infra.GetLogger().Printf("gRPC method: %s, request: %+v", info.FullMethod, req)

	// Call the gRPC handler to process the request
	resp, err = handler(ctx, req)

	// Log the outgoing response
	infra.GetLogger().Printf("gRPC method: %s, response: %+v", info.FullMethod, resp)

	return resp, err
}