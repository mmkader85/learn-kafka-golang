start-server:
	@echo "Starting HTTP server"
	@go run main.go httpserver

start-consumer:
	@echo "Starting Kafka consumer"
	@go run main.go consumer
