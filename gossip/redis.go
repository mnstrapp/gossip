package gossip

import (
	"context"
	"fmt"

	"github.com/mnstrapp/gossip/log"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func NewRedisClient(url string) (*redis.Client, error) {
	if redisClient != nil {
		return redisClient, nil
	}
	options, err := redis.ParseURL(url)
	if err != nil {
		log.Error(fmt.Errorf("parsing redis url: %s", err))
		return nil, err
	}
	redisClient = redis.NewClient(options)
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("failed to ping redis: %s", err)
		return nil, err
	}
	log.Info("redis connected")
	return redisClient, nil
}
