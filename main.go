package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/mnstrapp/gossip/gossip"
	"github.com/mnstrapp/gossip/log"
)

var (
	port     int
	redisUrl string
)

func init() {
	flag.IntVar(&port, "port", 3000, "Gossip port")
	flag.StringVar(&redisUrl, "redis-url", "redis://localhost:6379/0", "Redis url")
	flag.Parse()

	if p, ok := os.LookupEnv("PORT"); ok {
		port, _ = strconv.Atoi(p)
	}
	if port == 0 {
		port = getRandomPort()
	}

	if url, ok := os.LookupEnv("REDIS_URL"); ok {
		redisUrl = url
	}
}

func main() {
	if _, err := gossip.NewRedisClient(redisUrl); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	uri := fmt.Sprintf(":%d", port)
	server := gossip.NewServer(uri)
	log.Infof("server listening at %s", uri)
	if err := server.Listen(); err != nil {
		log.Errorf("server listen: %s", err)
		os.Exit(1)
	}
}

func getRandomPort() int {
	return rand.Intn(10000) + 40000
}
