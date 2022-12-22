package main

import (
	"context"
	"log"

	"github.com/TemurMannonov/medium_notification_service_with_kafka/config"
	"github.com/TemurMannonov/medium_notification_service_with_kafka/events"
)

func main() {
	cfg := config.Load(".")

	pubsub, err := events.New(cfg)
	if err != nil {
		log.Fatalf("failed to open pubsub %v", err)
	}

	pubsub.Run(context.Background())
}
