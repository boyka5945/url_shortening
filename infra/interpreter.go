package infra

import (
	"context"

	"google.golang.org/grpc"
)

func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// Log the incoming request
	GetLogger().Printf("gRPC method: %s, request: %+v", info.FullMethod, req)

	// Call the gRPC handler to process the request
	resp, err = handler(ctx, req)

	// Log the outgoing response
	GetLogger().Printf("gRPC method: %s, response: %+v", info.FullMethod, resp)

	return resp, err
}