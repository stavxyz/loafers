// main.go
package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

var Version string

func hello() (string, error) {
	return fmt.Sprintf("Hello ƛ! \n Version: %s", Version), nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
