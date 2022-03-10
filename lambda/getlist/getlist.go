package main

import (
	"encoding/json"
	"main/service"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func list() (Response, error) {
	res, _ := json.Marshal(service.Getlist())
	return Response{
		StatusCode: 200,
		Body:       string(res),
	}, nil
}

func main() {
	lambda.Start(list)
}
