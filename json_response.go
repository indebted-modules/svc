package svc

import (
	"github.com/aws/aws-lambda-go/events"
)

// Res creates a cors-enabled APIGatewayProxyResponse with JSON content type
func Res(statusCode int, body string) events.APIGatewayProxyResponse {
	return res(statusCode, body, "application/json")
}

// HTMLRes creates a cors-enabled APIGatewayProxyResponse with HTML content type
func HTMLRes(statusCode int, body string) events.APIGatewayProxyResponse {
	return res(statusCode, body, "text/html")
}

func res(statusCode int, body, contentType string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                contentType,
		},
		Body: body,
	}
}
