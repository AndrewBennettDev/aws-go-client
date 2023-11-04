package main

import(
  "os"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main(){
  region := os.Getenv("AWS_REGION")
  awsSession, err := session.NewSession(&aws.Config{
    Region: aws.String(region)},)

  if err != nil {
    return
  }

  dynaClient = dynamodb.New(awsSession)
  lambda.Start(handler)
}
