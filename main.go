package main

import (
	"fmt"
	"os"

	"go-lambda-slack/notification"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handler ...
func Handler() error {
	url := os.Getenv("WEB_HOOK_URL")
	err := notification.Send(url)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
