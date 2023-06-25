package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

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

	// // Initialize the DB
	// err = infra.InitDB()
	// if err != nil {
	// 	log.Fatalf("failed to initialize DB: %v", err)
	// }
	// defer infra.CloseDB()
	// infra.GetLogger().Printf("DB initialized")

	// // Initialize the Redis
	// err = infra.InitRedis()
	// if err != nil {
	// 	log.Fatalf("failed to initialize Redis: %v", err)
	// }
	// defer infra.CloseRedis()
	// infra.GetLogger().Printf("Redis initialized")

	// Register service in etcd
	var methodPaths []string
	for _, method := range pb.UrlShorteningService_ServiceDesc.Methods {
		methodPaths = append(methodPaths, method.MethodName)
	}
	var etcdEndpoints []string
	for _, v := range config.GetConfig().Etcd {
		etcdEndpoints = append(etcdEndpoints, fmt.Sprintf("%s:%d", v.Host, v.Port))
	}
	serviceRegistrar, err := infra.NewServiceRegistrar(context.Background(), etcdEndpoints, pb.UrlShorteningService_ServiceDesc.ServiceName, fmt.Sprintf("%s:%d", config.GetConfig().Server.Host, config.GetConfig().Server.Port), methodPaths, 10*time.Second)
	if err != nil {
		infra.GetLogger().Fatalf("failed to initialize service registrar: %v", err)
	}
	err = serviceRegistrar.RegisterService()
	if err != nil {
		infra.GetLogger().Fatalf("failed to register service: %v", err)
	}
	infra.GetLogger().Printf("service registered")

	// Start a background goroutine to periodically delete inactive services
	go func() {
		ticker := time.NewTicker(5 * time.Second) // Adjust the check interval as per your requirements
		defer ticker.Stop()

		for range ticker.C {
			err := serviceRegistrar.DeleteInactiveServices()
			if err != nil {
				log.Printf("Failed to delete inactive services: %v", err)
			}
		}
	}()

	go func () {
		serviceRegistrar.ComsumeLeaseKeepCh()
	}()

	// Initialize the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GetConfig().Server.Port))
	if err != nil {
		infra.GetLogger().Fatalf("failed to listen: %v", err)
	}

	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(infra.LoggingInterceptor),
	}
	server := grpc.NewServer(serverOptions...)

	pb.RegisterUrlShorteningServiceServer(server, service.NewUrlShorteningServiceServer())
	if err := server.Serve(lis); err != nil {
		infra.GetLogger().Fatalf("failed to serve: %v", err)
	}
}
