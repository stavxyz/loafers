// main.go
package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// Version is the git version of this lambda
var Version string

// Response is the lambda's return value data
type Response struct {
	LambdaContext   *lambdacontext.LambdaContext
	Version         string `json:"Version"`
	Message         string `json:"Message"`
	AwsRequestID    string `json:"AwsRequestID"`
	FunctionVersion string `json:"FunctionVersion"`
	FunctionName    string `json:"FunctionName"`
}

// Event is the data passed to the lambda function
type Event struct {
	Message string `json:"message"`
}

func hello(ctx context.Context, event Event) (string, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	res := &Response{
		LambdaContext: lc,
		Message:       event.Message,
		Version:       Version,
		// the following come from the aws lambda env
		AwsRequestID:    lc.AwsRequestID,
		FunctionVersion: lambdacontext.FunctionVersion,
		FunctionName:    lambdacontext.FunctionName,
	}
	content, err := json.Marshal(res)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
