package svc

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
)

// Event represents a envelop of a domain event
type Event struct {
	ID               string
	Type             string
	AggregateID      string
	AggregateType    string
	AggregateVersion int64
	Payload          string
	Created          time.Time
}

// EventMap maps event type to handler
type EventMap = map[string]eventHandler

// eventHandler handles events
type eventHandler = func(*Event) error

// EventHandler make a kinesis event handler specific to domain events
func EventHandler(eventMap EventMap) kinesisEventHandler {
	return func(kinesisEvent events.KinesisEvent) error {
		if len(kinesisEvent.Records) != 1 {
			err := fmt.Errorf("Unexpected KinesisEvent records length: %d. Expected count: %d", len(kinesisEvent.Records), 1)
			log.
				Error().
				Err(err).
				Msg(err.Error())
			return err
		}

		event := &Event{}
		err := json.Unmarshal(kinesisEvent.Records[0].Kinesis.Data, event)
		if err != nil {
			log.
				Error().
				Err(err).
				Msg("Failed parsing event")
			return err
		}

		handler, ok := eventMap[event.Type]
		if !ok {
			return nil
		}

		start := time.Now().UnixNano()

		err = handler(event)
		msg := "Consumed event"
		logContext := log.Info()

		if err != nil {
			msg = "Failed consuming event"
			logContext = log.Error().Err(err)
		}

		logContext.
			Str("EventID", event.ID).
			Str("EventType", event.Type).
			Str("AggregateID", event.AggregateID).
			Str("AggregateType", event.AggregateType).
			Int64("AggregateVersion", event.AggregateVersion).
			Time("Created", event.Created).
			Int64("DurationMS", (time.Now().UnixNano()-start)/int64(time.Millisecond)).
			Msg(msg)

		if err != nil {
			return err
		}

		return nil
	}
}
