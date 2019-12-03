package svc_test

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/indebted-modules/svc"
	"github.com/stretchr/testify/suite"
)

type MiddlewareSuite struct {
	suite.Suite
}

func TestMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareSuite))
}

func (s *MiddlewareSuite) TestExpectKinesisEventRecordsCallsThroughWhenCountIsExact() {
	var callThroughCount int
	callThroughHandler := func(event events.KinesisEvent) error {
		callThroughCount++
		return nil
	}

	expectOneHandler := svc.ExpectKinesisEventRecords(1, callThroughHandler)
	expectTwoHandler := svc.ExpectKinesisEventRecords(2, callThroughHandler)

	oneRecord := events.KinesisEvent{Records: []events.KinesisEventRecord{{}}}
	twoRecords := events.KinesisEvent{Records: []events.KinesisEventRecord{{}, {}}}

	err := expectOneHandler(oneRecord)
	s.Nil(err)
	s.Equal(1, callThroughCount)

	callThroughCount = 0
	err = expectTwoHandler(twoRecords)
	s.Equal(1, callThroughCount)
	s.Nil(err)
}

func (s *MiddlewareSuite) TestExpectKinesisEventRecordsErrorsWhenCountIsDifferent() {
	var callThroughCount int
	callThroughHandler := func(event events.KinesisEvent) error {
		callThroughCount++
		return nil
	}

	expectOneHandler := svc.ExpectKinesisEventRecords(1, callThroughHandler)
	expectTwoHandler := svc.ExpectKinesisEventRecords(2, callThroughHandler)

	zeroRecords := events.KinesisEvent{Records: []events.KinesisEventRecord{}}
	twoRecords := events.KinesisEvent{Records: []events.KinesisEventRecord{{}, {}}}
	threeRecords := events.KinesisEvent{Records: []events.KinesisEventRecord{{}, {}, {}}}

	err := expectOneHandler(zeroRecords)
	s.Equal("Unexpected KinesisEvent records length: 0. Expected count: 1", err.Error())
	s.Equal(0, callThroughCount)

	err = expectOneHandler(twoRecords)
	s.Equal("Unexpected KinesisEvent records length: 2. Expected count: 1", err.Error())
	s.Equal(0, callThroughCount)

	err = expectTwoHandler(zeroRecords)
	s.Equal("Unexpected KinesisEvent records length: 0. Expected count: 2", err.Error())
	s.Equal(0, callThroughCount)

	err = expectTwoHandler(threeRecords)
	s.Equal("Unexpected KinesisEvent records length: 3. Expected count: 2", err.Error())
	s.Equal(0, callThroughCount)
}

func (s *MiddlewareSuite) TestExpectDynamoDBEventRecordsCallsThroughWhenCountIsExact() {
	var callThroughCount int
	callThroughHandler := func(event events.DynamoDBEvent) error {
		callThroughCount++
		return nil
	}

	expectOneHandler := svc.ExpectDynamoDBEventRecords(1, callThroughHandler)
	expectTwoHandler := svc.ExpectDynamoDBEventRecords(2, callThroughHandler)

	oneRecord := events.DynamoDBEvent{Records: []events.DynamoDBEventRecord{{}}}
	twoRecords := events.DynamoDBEvent{Records: []events.DynamoDBEventRecord{{}, {}}}

	err := expectOneHandler(oneRecord)
	s.Nil(err)
	s.Equal(1, callThroughCount)

	callThroughCount = 0
	err = expectTwoHandler(twoRecords)
	s.Equal(1, callThroughCount)
	s.Nil(err)
}

func (s *MiddlewareSuite) TestExpectDynamoDBEventRecordsErrorsWhenCountIsDifferent() {
	var callThroughCount int
	callThroughHandler := func(event events.DynamoDBEvent) error {
		callThroughCount++
		return nil
	}

	expectOneHandler := svc.ExpectDynamoDBEventRecords(1, callThroughHandler)
	expectTwoHandler := svc.ExpectDynamoDBEventRecords(2, callThroughHandler)

	zeroRecords := events.DynamoDBEvent{Records: []events.DynamoDBEventRecord{}}
	twoRecords := events.DynamoDBEvent{Records: []events.DynamoDBEventRecord{{}, {}}}
	threeRecords := events.DynamoDBEvent{Records: []events.DynamoDBEventRecord{{}, {}, {}}}

	err := expectOneHandler(zeroRecords)
	s.Equal("Unexpected DynamoDBEvent records length: 0. Expected count: 1", err.Error())
	s.Equal(0, callThroughCount)

	err = expectOneHandler(twoRecords)
	s.Equal("Unexpected DynamoDBEvent records length: 2. Expected count: 1", err.Error())
	s.Equal(0, callThroughCount)

	err = expectTwoHandler(zeroRecords)
	s.Equal("Unexpected DynamoDBEvent records length: 0. Expected count: 2", err.Error())
	s.Equal(0, callThroughCount)

	err = expectTwoHandler(threeRecords)
	s.Equal("Unexpected DynamoDBEvent records length: 3. Expected count: 2", err.Error())
	s.Equal(0, callThroughCount)
}
