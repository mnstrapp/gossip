package gossip

import (
	"context"
	"net"

	"github.com/mnstrapp/gossip/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
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

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			p, _ := peer.FromContext(ss.Context())
			log.Infof("incoming stream request: %s->%s %s", p.Addr, p.LocalAddr, info.FullMethod)
			if err := handler(srv, ss); err != nil {
				log.Errorf("incoming stream error: %s->%s %s: %s", p.Addr, p.LocalAddr, info.FullMethod, err)
			}
			return err
		}),
		grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			p, _ := peer.FromContext(ctx)
			log.Infof("incoming unary request: %s->%s %s", p.Addr, p.LocalAddr, info.FullMethod)
			if resp, err = handler(ctx, req); err != nil {
				log.Errorf("incoming unary error: %s->%s %s: %s", p.Addr, p.LocalAddr, info.FullMethod, err)
			}
			return
		}),
	}

	server := grpc.NewServer(opts...)
	RegisterGossipApiServer(server, &service{})

	return server.Serve(lis)
}
