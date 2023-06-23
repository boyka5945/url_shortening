package test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"url_shortening/model"
)

func TestEnd2End(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := url.NewUrlShorteningServiceClient(conn)
	resp, err := client.CreateUrl(ctx, &url.CreateUrlRequest{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}