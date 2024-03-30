package gossip

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/mnstrapp/gossip/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var streams = make(map[string]GossipApi_SubscribeToEventsServer)

type service struct {
	UnimplementedGossipApiServer
}

func (s *service) SubscribeToEvents(user *User, stream GossipApi_SubscribeToEventsServer) error {
	id := user.GetId()
	if _, ok := streams[id]; !ok {
		streams[id] = stream
	}
	log.Info(fmt.Sprintf("%s subscribed to events", user.Id))
	broadcastEvent(userJoinedEvent(user))
	return nil
}

func (s *service) SendEvent(ctx context.Context, event *Event) (*emptypb.Empty, error) {
	if event == nil {
		err := fmt.Errorf("empty event")
		log.Error(err)
		return nil, err
	}
	broadcastEvent(event)
	return nil, nil
}

func (s *service) UnsubscribeFromEvents(ctx context.Context, user *User) (*emptypb.Empty, error) {
	if user == nil {
		err := fmt.Errorf("empty user")
		log.Error(err)
		return nil, err
	}
	broadcastEvent(userLeftEvent(user.Id))
	return nil, nil
}

func broadcastEvent(event *Event) {
	if event.ToId != nil {
		dmEvent(event)
		return
	}
	for id, stream := range streams {
		if event.FromId != nil && id == *event.FromId {
			continue
		}
		go func() {
			if err := stream.Send(event); err != nil {
				log.Error(err)
			}
		}()
	}
}

func dmEvent(event *Event) {
	stream, ok := streams[*event.FromId]
	if !ok {
		log.Info(fmt.Sprintf("user %s left", *event.FromId))
		dmEvent(userLeftEvent(*event.ToId))
		return
	}
	if err := stream.Send(event); err != nil {
		log.Error(err)
	}
}

func userJoinedEvent(user *User) *Event {
	return &Event{
		Type:      EventType_UserEventType,
		FromId:    &user.Id,
		Timestamp: uint64(time.Now().Unix()),
		Body: &Event_User{
			User: &UserEvent{
				UserId: user.Id,
				Status: NetworkStatus_Online,
			},
		},
	}
}

func userLeftEvent(id string) *Event {
	return &Event{
		Type:      EventType_UserEventType,
		FromId:    &id,
		Timestamp: uint64(time.Now().Unix()),
		Body: &Event_User{
			User: &UserEvent{
				UserId: id,
				Status: NetworkStatus_Offline,
			},
		},
	}
}

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
