package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/mnstrapp/gossip/gossip"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 3000, "Gossip port")
	flag.Parse()

	if p, ok := os.LookupEnv("PORT"); ok {
		port, _ = strconv.Atoi(p)
	}
	if port == 0 {
		port = getRandomPort()
	}
}

func main() {
	uri := fmt.Sprintf(":%d", port)
	server := gossip.NewServer(uri)
	fmt.Fprintf(os.Stdout, "[INFO] server listening at %s\n", uri)
	if err := server.Listen(); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] server listen: %s\n", err)
		os.Exit(1)
	}
}

func getRandomPort() int {
	return rand.Intn(10000) + 40000
}
