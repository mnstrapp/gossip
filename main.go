package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/mnstrapp/gossip/gossip"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "Gossip host")
	flag.IntVar(&port, "port", 0, "Gossip port")
	flag.Parse()

	if port == 0 {
		port = getRandomPort()
	}
}

func main() {
	uri := fmt.Sprintf("%s:%d", host, port)
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
