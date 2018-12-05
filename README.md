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
./build.sh
```

This will produce a `main.zip` file to upload to aws lambda (or s3).

Your value for `handler` will be simply `main`.

More info on the lambda execution env:
https://docs.aws.amazon.com/lambda/latest/dg/current-supported-versions.html

