package main

import (
	"fmt"
	"log"
	"net"
	"url_shortening/config"
	"url_shortening/infra"
	pb "url_shortening/model"
	"url_shortening/service"

	"google.golang.org/grpc"
)

func main() {
	// Initialize the logger
	err := infra.InitLog()
	if err != nil {
		log.Fatalf("failed to initialize log: %v", err)
	}
	infra.GetLogger().Printf("log initialized")

	// Initialize the config
	err = config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	infra.GetLogger().Printf("config loaded, config: %+v", config.GetConfig())

	// Initialize the DB
	err = infra.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize DB: %v", err)
	}
	defer infra.CloseDB()
	infra.GetLogger().Printf("DB initialized")

	// Initialize the Redis
	err = infra.InitRedis()
	if err != nil {
		log.Fatalf("failed to initialize Redis: %v", err)
	}
	defer infra.CloseRedis()
	infra.GetLogger().Printf("Redis initialized")

	// Initialize the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GetConfig().Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(infra.LoggingInterceptor),
	}
	server := grpc.NewServer(serverOptions...)

	pb.RegisterUrlShorteningServiceServer(server, service.NewUserServiceServer())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
