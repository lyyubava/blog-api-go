package notify

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

type EventInfo struct {
	EventName string
	EventTime time.Time
}
type Event struct {
	EventDetails string
	EventInfo    EventInfo
}

func Publish(event Event) error {
	eventMsg, _ := json.Marshal(&event)
	context := context.Background()
	client, _ := pubsub.NewClient(context, os.Getenv("PROJECT_ID"))
	topic := client.Topic(os.Getenv("PUBSUB_TOPIC"))
	publishResult := topic.Publish(context, &pubsub.Message{
		Data: []byte(eventMsg),
	})
	_, err := publishResult.Get(context)
	if err != nil {
		return err
	}

	return nil
}
