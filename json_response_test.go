package svc_test

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/indebted-modules/svc"
	"github.com/stretchr/testify/suite"
)

type JSONResponseSuite struct {
	suite.Suite
}

func TestJSONResponseSuite(t *testing.T) {
	suite.Run(t, new(JSONResponseSuite))
}

func (s *JSONResponseSuite) TestRes() {
	expected := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
		Body: "foo",
	}
	s.Equal(expected, svc.Res(http.StatusOK, "foo"))
}

func (s *JSONResponseSuite) TestHTMLRes() {
	expected := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "text/html",
		},
		Body: "bar",
	}
	s.Equal(expected, svc.HTMLRes(http.StatusOK, "bar"))
}
