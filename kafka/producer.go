package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hamba/avro/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

type Order struct {
	OrderID   string `json:"order_id"`
	UserID    string `json:"user_id"`
	MenuID    string `json:"menu_id"`
	Quantity  int    `json:"quantity"`
	OrderTime string `json:"order_time"`
}

func ProduceMessage(message []byte) {
	/*w := kafka.NewWriter(kafka.WriterConfig{
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
	}*/

	/*orderSchemaStr, err := os.ReadFile("schemas/order-schema.avsc")
	if err != nil {
		log.Fatal(err)
	}*/

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
	})
	if err != nil {
		log.Fatalf("Unable to connect to kafka server: %v\n", err)
	}
	defer p.Close()

	orderSchema, err := avro.ParseFiles("schemas/order-schema.avsc")
	if err != nil {
		log.Fatalf("Unable to parse avro order schema: %v\n", err)
	}

	orderData, err := avro.Marshal(orderSchema, message)
	if err != nil {
		log.Fatalf("Message doesn't match the schema: %v\n%v\n", orderSchema, message)
	}

	topic := os.Getenv("KAFKA_TOPIC")
	err = p.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: orderData,
		}, nil)
	if err != nil {
		log.Fatalf("Unable to produce message: %v\n", err)
	}

	log.Printf("Event produced successfully: %v\n", message)
}
