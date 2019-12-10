package svc_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/indebted-modules/svc"
	"github.com/stretchr/testify/suite"
)

type EventHandlerSuite struct {
	suite.Suite
}

func TestEventHandlerSuite(t *testing.T) {
	suite.Run(t, new(EventHandlerSuite))
}

func (s *EventHandlerSuite) TestMappedEvents() {
	var eventReceived *svc.Event
	handler := svc.EventHandler(svc.EventMap{"Foo": func(event *svc.Event) error {
		eventReceived = event
		return nil
	}})

	event := &svc.Event{
		ID:               "123",
		Type:             "Foo",
		AggregateID:      "abc",
		AggregateType:    "Bar",
		AggregateVersion: 1,
		Payload:          "{}",
		Created:          time.Unix(0, 0),
	}

	b, err := json.Marshal(event)
	s.Nil(err)

	kinesisEvent := events.KinesisEvent{Records: []events.KinesisEventRecord{
		{Kinesis: events.KinesisRecord{Data: b}},
	}}

	err = handler(kinesisEvent)
	s.Nil(err)
	s.Equal(event, eventReceived)
}

func (s *EventHandlerSuite) TestUnmappedEvents() {
	invoked := false
	handler := svc.EventHandler(svc.EventMap{"Unmapped": func(event *svc.Event) error {
		invoked = true
		return nil
	}})

	event := &svc.Event{
		ID:               "123",
		Type:             "Foo",
		AggregateID:      "abc",
		AggregateType:    "Bar",
		AggregateVersion: 1,
		Payload:          "{}",
		Created:          time.Unix(0, 0),
	}

	b, err := json.Marshal(event)
	s.Nil(err)

	kinesisEvent := events.KinesisEvent{Records: []events.KinesisEventRecord{
		{Kinesis: events.KinesisRecord{Data: b}},
	}}

	err = handler(kinesisEvent)
	s.Nil(err)
	s.False(invoked)
}

func (s *EventHandlerSuite) TestUnmarshallingError() {
	invoked := false
	handler := svc.EventHandler(svc.EventMap{"Foo": func(event *svc.Event) error {
		invoked = true
		return nil
	}})

	kinesisEvent := events.KinesisEvent{Records: []events.KinesisEventRecord{
		{Kinesis: events.KinesisRecord{Data: []byte("invalid-json")}},
	}}

	err := handler(kinesisEvent)
	s.Error(err)
	s.False(invoked)
}

func (s *EventHandlerSuite) TestExpectedRecordCount() {
	invoked := false
	handler := svc.EventHandler(svc.EventMap{"Foo": func(event *svc.Event) error {
		invoked = true
		return nil
	}})

	event := &svc.Event{
		ID:               "123",
		Type:             "Foo",
		AggregateID:      "abc",
		AggregateType:    "Bar",
		AggregateVersion: 1,
		Payload:          "{}",
		Created:          time.Unix(0, 0),
	}

	b, err := json.Marshal(event)
	s.Nil(err)

	kinesisEvent := events.KinesisEvent{Records: []events.KinesisEventRecord{
		{Kinesis: events.KinesisRecord{Data: b}},
	}}

	err = handler(kinesisEvent)
	s.Nil(err)
	s.True(invoked)
}

func (s *EventHandlerSuite) TestUnexpectedRecordCount() {
	handler := svc.EventHandler(svc.EventMap{"Foo": func(event *svc.Event) error {
		return nil
	}})

	event := &svc.Event{
		ID:               "123",
		Type:             "Foo",
		AggregateID:      "abc",
		AggregateType:    "Bar",
		AggregateVersion: 1,
		Payload:          "{}",
		Created:          time.Unix(0, 0),
	}

	b, err := json.Marshal(event)
	s.Nil(err)

	kinesisEvent := events.KinesisEvent{Records: []events.KinesisEventRecord{
		{Kinesis: events.KinesisRecord{Data: b}},
		{Kinesis: events.KinesisRecord{Data: b}},
	}}

	err = handler(kinesisEvent)
	s.Error(err)
}
