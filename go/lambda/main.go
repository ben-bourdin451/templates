package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"function/src"
)

func application(ctx context.Context, payload src.Event) error {
	err := src.HandleEvent(payload)

	return handleError(err)
}

func handleError(err error) error {
	log.Println("An error occured", err)
	return err
}

// Lambda entrypoint
func main() {
	lambda.Start(application)
}
