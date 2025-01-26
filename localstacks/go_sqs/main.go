package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqsTypes "github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const (
	queueName = "test-queue"
)

func main() {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			if service == sqs.ServiceID {
				return aws.Endpoint{URL: "http://localhost:4566"}, nil
			}
			return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
		})),
	)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Create SQS client
	sqsClient := sqs.NewFromConfig(cfg)

	// Create SQS queue
	queueURL, err := createQueue(sqsClient, queueName)
	if err != nil {
		log.Fatalf("failed to create queue: %v", err)
	}
	fmt.Printf("Queue URL: %s\n", queueURL)

	// Send message to the queue
	messageBody := "Hello, SQS!"
	if err := sendMessage(sqsClient, queueURL, messageBody); err != nil {
		log.Fatalf("failed to send message: %v", err)
	}
	fmt.Println("Message sent successfully")

	// Receive messages from the queue
	messages, err := receiveMessages(sqsClient, queueURL)
	if err != nil {
		log.Fatalf("failed to receive messages: %v", err)
	}
	for _, msg := range messages {
		fmt.Printf("Received message: %s\n", aws.ToString(msg.Body))

		// Delete the message after processing
		if err := deleteMessage(sqsClient, queueURL, msg.ReceiptHandle); err != nil {
			log.Printf("failed to delete message: %v", err)
		} else {
			fmt.Println("Message deleted successfully")
		}
	}
}

func createQueue(client *sqs.Client, name string) (string, error) {
	output, err := client.CreateQueue(context.TODO(), &sqs.CreateQueueInput{
		QueueName: &name,
	})
	if err != nil {
		return "", err
	}
	return aws.ToString(output.QueueUrl), nil
}

func sendMessage(client *sqs.Client, queueURL, messageBody string) error {
	_, err := client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &queueURL,
		MessageBody: &messageBody,
	})
	return err
}

func receiveMessages(client *sqs.Client, queueURL string) ([]sqsTypes.Message, error) {
	output, err := client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            &queueURL,
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     5,
	})
	if err != nil {
		return nil, err
	}
	return output.Messages, nil
}

func deleteMessage(client *sqs.Client, queueURL string, receiptHandle *string) error {
	_, err := client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
		QueueUrl:      &queueURL,
		ReceiptHandle: receiptHandle,
	})
	return err
}
