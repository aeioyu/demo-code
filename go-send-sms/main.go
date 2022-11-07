package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func main() {
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     "",
				SecretAccessKey: "",
				SessionToken:    "",
				Source:          "",
			},
		}))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sns.NewFromConfig(config)

	params := &sns.PublishInput{
		Message:     aws.String("this is test sms"),
		PhoneNumber: aws.String("+66yournumber"),
	}
	resp, err := client.Publish(context.TODO(), params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
