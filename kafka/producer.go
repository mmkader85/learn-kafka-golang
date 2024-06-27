package kafka

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func ProduceMessage(message []byte) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{os.Getenv("KAFKA_BROKER")},
		Topic:   os.Getenv("KAFKA_TOPIC"),
	})

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: message,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
