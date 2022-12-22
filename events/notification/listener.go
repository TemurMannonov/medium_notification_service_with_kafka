package notification

import (
	"context"
	"encoding/json"

	emailPkg "github.com/TemurMannonov/medium_notification_service_with_kafka/pkg/email"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type Notification struct {
	To      string
	Type    string
	Body    map[string]string
	Subject string
}

func (c *NotificationService) SendEmail(ctx context.Context, event cloudevents.Event) error {
	var (
		req Notification
	)

	err := json.Unmarshal(event.DataEncoded, &req)
	if err != nil {
		return err
	}

	err = emailPkg.SendEmail(&c.cfg, &emailPkg.SendEmailRequest{
		To:      []string{req.To},
		Subject: req.Subject,
		Body:    req.Body,
		Type:    req.Type,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *NotificationService) SendSMS(ctx context.Context, event cloudevents.Event) error {

	return nil
}
