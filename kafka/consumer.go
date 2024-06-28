package kafka

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func init() {
	godotenv.Load()
}

func StartConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_BROKER")},
		Topic:   os.Getenv("KAFKA_TOPIC"),
		GroupID: "consumer-group-id",
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("received: %s", string(msg.Value))
	}
}
