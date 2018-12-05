# loafers

Lambda meets gophers.

### start

Install `aws-lambda-go` project

> Libraries, samples and tools to help Go developers develop AWS Lambda
> functions.

```
go get github.com/aws/aws-lambda-go/lambda
``` 

### build

This builds to run in the aws lambda runtime.

```
GOOS=linux GOARCH=amd64 go build -o main main.go
```

More info on the lambda execution env:
https://docs.aws.amazon.com/lambda/latest/dg/current-supported-versions.html
