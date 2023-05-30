package notify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type EventInfo struct {
	EventName      string
	EventTime      time.Time
	EventUser      string
	EventUserEmail string
}
type Event struct {
	EventDetails string
	EventInfo    EventInfo
}

func Publish(event Event) error {
	eventMsg, _ := json.Marshal(&event)
	context := context.Background()
	creds, _ := base64.StdEncoding.DecodeString(os.Getenv("SERVICE_ACCOUNT_KEY"))
	client, _ := pubsub.NewClient(context, os.Getenv("PROJECT_ID"), option.WithCredentialsJSON(creds))
	topic := client.Topic(os.Getenv("PUBSUB_TOPIC"))
	publishResult := topic.Publish(context, &pubsub.Message{
		Data: []byte(eventMsg),
	})

	_, err := publishResult.Get(context)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
