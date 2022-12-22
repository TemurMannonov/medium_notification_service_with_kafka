package notification

import (
	event "github.com/TemurMannonov/medium_notification_service_with_kafka/pkg/messagebroker"

	"github.com/TemurMannonov/medium_notification_service_with_kafka/config"
)

type NotificationService struct {
	cfg   config.Config
	kafka *event.Kafka
}

func New(cfg config.Config, kafka *event.Kafka) *NotificationService {
	return &NotificationService{
		cfg:   cfg,
		kafka: kafka,
	}
}

func (c *NotificationService) RegisterConsumers() {
	notificationRoute := "v1.notification_service.send_email"

	c.kafka.AddConsumer(
		notificationRoute, // consumer name
		notificationRoute, // topic
		c.SendEmail,       // handlerFunction
	)

	// notificationRouteSendSMS := "v1.notification_service.send_sms"

	// c.kafka.AddConsumer(
	// 	notificationRouteSendSMS, // consumer name
	// 	notificationRouteSendSMS, // topic
	// 	c.SendSMS,                // handlerFunction
	// )
}
