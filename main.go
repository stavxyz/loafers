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
	*lambdacontext.LambdaContext
	Version         string `json:"Version"`
	Message         string `json:"Message"`
	FunctionVersion string `json:"FunctionVersion"`
	FunctionName    string `json:"FunctionName"`
}

type Event struct {
	Message string `json:"message"`
}

func hello(ctx context.Context, event Event) (string, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	res := &Response{
		LambdaContext: lc,
		Message:       event.Message,
		Version:       Version,
		// The following are globals in the lambda env
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
