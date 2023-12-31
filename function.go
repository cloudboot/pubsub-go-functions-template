package cloudboot

import (
  "context"
  "fmt"
  "log"

  "github.com/GoogleCloudPlatform/functions-framework-go/functions"
  "github.com/cloudevents/sdk-go/v2/event"
)

func init() {
  functions.CloudEvent("Main", main)
}

// MessagePublishedData contains the full Pub/Sub message
// See the documentation for more details:
// https://cloud.google.com/eventarc/docs/cloudevents#pubsub
type MessagePublishedData struct {
  Message PubSubMessage
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
  Data []byte `json:"data"`
}

// helloPubSub consumes a CloudEvent message and extracts the Pub/Sub message.
func main(ctx context.Context, e event.Event) error {
  var msg MessagePublishedData
  if err := e.DataAs(&msg); err != nil {
    return fmt.Errorf("event.DataAs: %v", err)
  }

  name := string(msg.Message.Data) // Automatically decoded from base64.
  log.Printf("It works %s!", name)
  return nil
}
