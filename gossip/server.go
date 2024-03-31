package gossip

import (
	"net"

	"github.com/mnstrapp/gossip/log"
	"google.golang.org/grpc"
)

type Server struct {
	Uri string
}

func NewServer(uri string) *Server {
	return &Server{uri}
}

func (s *Server) Listen() error {
	lis, err := net.Listen("tcp", s.Uri)
	if err != nil {
		log.Error(err)
		return err
	}

	server := grpc.NewServer()
	RegisterGossipApiServer(server, &service{})

	return server.Serve(lis)
}
