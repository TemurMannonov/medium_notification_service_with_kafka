package event

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/TemurMannonov/medium_notification_service_with_kafka/config"

	"github.com/Shopify/sarama"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type Kafka struct {
	cfg          config.Config
	consumers    map[string]*Consumer
	saramaConfig *sarama.Config
}

func NewKafka(cfg config.Config) (*Kafka, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0

	kafka := &Kafka{
		cfg:          cfg,
		consumers:    make(map[string]*Consumer),
		saramaConfig: saramaConfig,
	}

	return kafka, nil
}

func (r *Kafka) RunConsumers(ctx context.Context) {
	var wg sync.WaitGroup

	for _, consumer := range r.consumers {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *Consumer) {
			defer wg.Done()

			err := c.cloudEventClient.StartReceiver(context.Background(), func(ctx context.Context, event cloudevents.Event) {
				err := c.handler(ctx, event)
				if err != nil {
					fmt.Println(err)
				}
			})

			log.Panic("Failed to start consumer", err)
		}(&wg, consumer)
		fmt.Println("Key:", consumer.topic, "=>", "consumer:", consumer)
	}
	wg.Wait()
}
