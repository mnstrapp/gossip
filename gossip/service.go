package gossip

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mnstrapp/gossip/log"
)

type service struct {
	UnimplementedGossipApiServer
}

func (s *service) SubscribeToEvents(user *User, stream GossipApi_SubscribeToEventsServer) error {
	log.Infof("%s subscribed to events", user.Id)
	broadcastEvent(userJoinedEvent(user))

	ctx := context.Background()
	pubsub := redisClient.Subscribe(ctx, "events")
	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Errorf("receiving from redis: %s", err)
	}
	subChan := pubsub.Channel()

	for {
		redisMessage := <-subChan
		eventMessage := []byte(redisMessage.Payload)

		var event Event
		if err := json.Unmarshal(eventMessage, &event); err != nil {
			log.Errorf("unmarshaling event message: %s", err)
		}

		if event.Type == EventType_DoneEventType && event.GetDone().UserId == user.Id {
			broadcastEvent(userLeftEvent(user.Id))
			pubsub.Close()
			return nil
		}
		log.Infof("event.FromId: %s, user.Id: %s", *event.FromId, user.Id)
		if event.ToId != nil && *event.ToId != user.Id {
			continue
		}
		if event.FromId != nil && *event.FromId == user.Id {
			continue
		}
		if err := stream.Send(&event); err != nil {
			log.Errorf("sending event: %s", err)
		}
	}
}

func (s *service) SendEvent(ctx context.Context, event *Event) (*Empty, error) {
	if event == nil {
		err := fmt.Errorf("empty event")
		log.Error(err)
		return nil, err
	}
	broadcastEvent(event)
	return nil, nil
}

func (s *service) UnsubscribeFromEvents(ctx context.Context, user *User) (*Empty, error) {
	if user == nil {
		err := fmt.Errorf("empty user")
		log.Error(err)
		return nil, err
	}
	log.Infof("%s unsubscribed from events", user.Id)
	broadcastEvent(doneEvent(user.Id))
	return nil, nil
}

func broadcastEvent(event *Event) {
	eventMessage, err := json.Marshal(event)
	if err != nil {
		log.Errorf("broadcasting: marshaling event message: %s", err)
		return
	}
	if err := redisClient.Publish(context.Background(), "events", eventMessage).Err(); err != nil {
		log.Errorf("broadcasting: publishing event: %s", err)
	}
}

func userJoinedEvent(user *User) *Event {
	return &Event{
		Type:      EventType_UserEventType,
		FromId:    &user.Id,
		Timestamp: uint64(time.Now().Unix()),
		User: &UserEvent{
			UserId: user.Id,
			Status: NetworkStatus_Online,
		},
	}
}

func userLeftEvent(id string) *Event {
	return &Event{
		Type:      EventType_UserEventType,
		FromId:    &id,
		Timestamp: uint64(time.Now().Unix()),
		User: &UserEvent{
			UserId: id,
			Status: NetworkStatus_Offline,
		},
	}
}

func doneEvent(id string) *Event {
	return &Event{
		Type:      EventType_DoneEventType,
		FromId:    &id,
		Timestamp: uint64(time.Now().Unix()),
		Done: &DoneEvent{
			UserId: id,
			Status: NetworkStatus_Offline,
		},
	}
}
