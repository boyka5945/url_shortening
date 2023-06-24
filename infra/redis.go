package infra

import (
	"context"
	"fmt"
	"url_shortening/config"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.ClusterClient
)

func RedisClient() *redis.ClusterClient {
	return redisClient
}

func InitRedis() error {
    // Initialize Redis Cluster client
	redisConfigs := config.GetConfig().Redis
	var redisAddrs []string
	for _, redisConfig := range redisConfigs {
		redisAddrs = append(redisAddrs, fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port))
	}
    redisClient = redis.NewClusterClient(&redis.ClusterOptions{
        Addrs: redisAddrs,
        Password: "",
    })

    // Test the Redis Cluster connection
    pong, err := redisClient.Ping(context.Background()).Result()
    if err != nil {
        GetLogger().Fatalf("Failed to connect to Redis Cluster: %v", err)
    }
    GetLogger().Println("Connected to Redis Cluster:", pong)
	return nil
}

func CloseRedis() {
	// Close the Redis Cluster client when done
	err := redisClient.Close()
	if err != nil {
		GetLogger().Fatalf("Failed to close Redis Cluster connection: %v", err)
	}
}
