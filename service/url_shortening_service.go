package service

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"url_shortening/infra"
	pb "url_shortening/model"
	model "url_shortening/model/models"
)

func NewUrlShorteningServiceServer() pb.UrlShorteningServiceServer {
	return &UrlShorteingService{}
}

// Implement the UserServiceServer interface
type UrlShorteingService struct {
	pb.UnimplementedUrlShorteningServiceServer
}

func (s *UrlShorteingService) CreateUrl(ctx context.Context, req *pb.CreateUrlRequest) (*pb.CreateUrlResponse, error) {
	if req.GetOriginalUrl() == "" {
		return nil, fmt.Errorf("original url is empty")
	}
	if req.GetApiDevKey() == "" {
		return nil, fmt.Errorf("api dev key is empty")
	}
	
	if req.GetCustomShortUrl() != "" {
		// check if the custom short url is already used
		d := infra.MasterSession(ctx).Where("short_url = ?", req.GetCustomShortUrl()).Find(&model.UrlTable{})
		if d.Error != nil {
			return nil, fmt.Errorf("failed to query db: %v", d.Error)
		}
		if d.RowsAffected > 0 {
			return nil, fmt.Errorf("custom short url is already used")
		}
		_, err := s.StoreUrlTable(ctx, req.GetOriginalUrl(), req.GetCustomShortUrl())
		if err != nil {
			return nil, fmt.Errorf("failed to store url mapping: %v", err)
		}
		return &pb.CreateUrlResponse{
			ShortUrl: req.GetCustomShortUrl(),
		}, nil
	}

	// generate a short url using base62 encoding
	shortUrl, err := s.GenerateShortUrl(ctx, req.GetOriginalUrl())
	if err != nil {
		return nil, fmt.Errorf("failed to generate short url: %v", err)
	}

	return &pb.CreateUrlResponse{
		ShortUrl: shortUrl,
	}, nil
}

func (s *UrlShorteingService) AccessUrl(ctx context.Context, req *pb.AccessUrlRequest) (*pb.AccessUrlResponse, error) {
	// get the original url from db and redirect 302 to it 
	originalUrl, err := s.GetOriginalUrl(ctx, req.GetShortUrl())
	if err != nil {
		return nil, fmt.Errorf("failed to get original url: %v", err)
	}
	return &pb.AccessUrlResponse{
		OriginalUrl: originalUrl,
	}, nil
}

func (s *UrlShorteingService) GenerateShortUrl(ctx context.Context, originalUrl string) (string, error) {
    // generate a short url using base62 encoding
    const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    
    urlTable, err := s.StoreUrlTable(ctx, originalUrl, "")
    if err != nil {
        return "", fmt.Errorf("failed to store url mapping: %v", err)
    }

    id := urlTable.ID
    var result strings.Builder
    for id > 0 {
        result.WriteByte(base62[id % 62])
        id /= 62
    }

    shortUrl := result.String()
    err = s.UpdateUrlTable(ctx, urlTable.ID, shortUrl)
    if err != nil {
        return "", fmt.Errorf("failed to update url mapping: %v", err)
    }

    return shortUrl, nil
}


func (s *UrlShorteingService) StoreUrlTable(ctx context.Context, originalUrl string, customShortUrl string) (*model.UrlTable, error) {
	url := &model.UrlTable{
		OriginalUrl: originalUrl,
		ShortUrl: customShortUrl,
		Offset: int64(rand.Intn(62)),
	}
	d := infra.MasterSession(ctx).Create(url)
	if d.Error != nil {
		return nil, fmt.Errorf("failed to store url mapping in db: %v", d.Error)
	}
	return url, nil
}

func (s *UrlShorteingService) UpdateUrlTable(ctx context.Context, id int64, shortUrl string) error {
    return infra.MasterSession(ctx).Model(&model.UrlTable{}).Where("id = ?", id).Update("short_url", shortUrl).Error
}

func (s *UrlShorteingService) GetOriginalUrl(ctx context.Context, shortUrl string) (string, error) {
	// get the original url from db
	var urlTable model.UrlTable
	d := infra.MasterSession(ctx).Where("short_url = ?", shortUrl).Find(&urlTable)
	if d.Error != nil {
		return "", fmt.Errorf("failed to query db: %v", d.Error)
	}
	if d.RowsAffected == 0 {
		return "", fmt.Errorf("short url is not found")
	}
	return urlTable.OriginalUrl, nil
}