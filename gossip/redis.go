package gossip

import (
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
	return redisClient, nil
}
