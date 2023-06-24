package service

import (
	"context"
	pb "url_shortening/model"
)

func NewUrlShorteningServiceServer() pb.UrlShorteningServiceServer {
	return &UrlShorteingService{}
}

// Implement the UserServiceServer interface
type UrlShorteingService struct {
	pb.UnimplementedUrlShorteningServiceServer
}

func (s *UrlShorteingService) CreateUrl(ctx context.Context, req *pb.CreateUrlRequest) (*pb.CreateUrlResponse, error) {
	return &pb.CreateUrlResponse{
		ShortUrl: "http://www.google.com",
  }, nil
}

func (s *UrlShorteingService) AccessUrl(ctx context.Context, req *pb.AccessUrlRequest) (*pb.AccessUrlResponse, error) {
	return &pb.AccessUrlResponse{
		OriginalUrl: "http://www.google.com",
  }, nil
}
