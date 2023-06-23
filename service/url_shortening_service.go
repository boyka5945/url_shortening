package service

import (
	"context"
	pb "url_shortening/model"
)

func NewUserServiceServer() pb.UrlShorteningServiceServer {
	return &UrlShorteingServicec{}
}

// Implement the UserServiceServer interface
type UrlShorteingServicec struct {
	pb.UnimplementedUrlShorteningServiceServer
}

func (s *UrlShorteingServicec) CreateUrl(ctx context.Context, req *pb.CreateUrlRequest) (*pb.CreateUrlResponse, error) {
	return &pb.CreateUrlResponse{
		ShortUrl: "http://www.google.com",
  }, nil
}

func (s *UrlShorteingServicec) AccessUrl(ctx context.Context, req *pb.AccessUrlRequest) (*pb.AccessUrlResponse, error) {
	return &pb.AccessUrlResponse{
		OriginalUrl: "http://www.google.com",
  }, nil
}
