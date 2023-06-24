package test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"url_shortening/model"
)

func TestEnd2End(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serviceHost, err := resolveServiceAddressFromEtcd("localhost:2379", "url.UrlShorteningService", "CreateUrl")
	if err != nil {
		t.Fatalf("failed to resolve service address from etcd: %v", err)
	}

	conn, err := grpc.DialContext(ctx, serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := url.NewUrlShorteningServiceClient(conn)
	resp, err := client.CreateUrl(ctx, &url.CreateUrlRequest{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func resolveServiceAddressFromEtcd(etcdEndpoints, serviceName, methodPath string) (string, error) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(etcdEndpoints, ","),
	})
	if err != nil {
		return "", fmt.Errorf("failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()

	key := fmt.Sprintf("/service/%s/%s", serviceName, methodPath)

	resp, err := etcdClient.Get(context.Background(), key)
	if err != nil {
		return "", fmt.Errorf("failed to resolve service address from etcd: %v", err)
	}

	if len(resp.Kvs) == 0 {
		return "", fmt.Errorf("no service address found in etcd")
	}

	return string(resp.Kvs[0].Value), nil
}
