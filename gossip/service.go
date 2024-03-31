package gossip

import (
	context "context"
	"fmt"
	"time"

	"github.com/mnstrapp/gossip/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var streams = make(map[string]GossipApi_SubscribeToEventsServer)
var eventChan = make(chan *Event)
var doneChan = make(chan string)

type service struct {
	UnimplementedGossipApiServer
	Context context.Context
}

func (s *service) SubscribeToEvents(user *User, stream GossipApi_SubscribeToEventsServer) error {
	id := user.GetId()
	if _, ok := streams[id]; !ok {
		streams[id] = stream
	}
	log.Info(fmt.Sprintf("%s subscribed to events", user.Id))
	broadcastEvent(userJoinedEvent(user))

	for {
		select {
		case event := <-eventChan:
			broadcastEvent(event)
			continue
		case id := <-doneChan:
			if id == user.Id {
				return nil
			}
			continue
		default:
			continue
		}
	}
}

func (s *service) SendEvent(ctx context.Context, event *Event) (*emptypb.Empty, error) {
	if event == nil {
		err := fmt.Errorf("empty event")
		log.Error(err)
		return nil, err
	}
	eventChan <- event
	return nil, nil
}

func (s *service) UnsubscribeFromEvents(ctx context.Context, user *User) (*emptypb.Empty, error) {
	if user == nil {
		err := fmt.Errorf("empty user")
		log.Error(err)
		return nil, err
	}
	log.Info(fmt.Sprintf("%s unsubscribed from events", user.Id))
	eventChan <- userLeftEvent(user.Id)
	if _, ok := streams[user.Id]; ok {
		doneChan <- user.Id
		delete(streams, user.Id)
	}
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
		if err := stream.Send(event); err != nil {
			log.Error(fmt.Errorf("error sending broadcast: %s", err))
		}
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
