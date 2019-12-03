package svc

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
)

type kinesisEventHandler = func(events.KinesisEvent) error
type dynamodbEventHandler = func(events.DynamoDBEvent) error

// ExpectKinesisEventRecords validates the number of records in the event
func ExpectKinesisEventRecords(count int, fn kinesisEventHandler) kinesisEventHandler {
	return func(event events.KinesisEvent) error {
		records := len(event.Records)
		if records != count {
			err := fmt.Errorf("Unexpected KinesisEvent records length: %d. Expected count: %d", records, count)
			log.
				Error().
				Err(err).
				Msg(err.Error())
			return err
		}
		return fn(event)
	}
}

// ExpectDynamoDBEventRecords validates the number of records in the event
func ExpectDynamoDBEventRecords(count int, fn dynamodbEventHandler) dynamodbEventHandler {
	return func(event events.DynamoDBEvent) error {
		records := len(event.Records)
		if records != count {
			err := fmt.Errorf("Unexpected DynamoDBEvent records length: %d. Expected count: %d", records, count)
			log.
				Error().
				Err(err).
				Msg(err.Error())
			return err
		}
		return fn(event)
	}
}
