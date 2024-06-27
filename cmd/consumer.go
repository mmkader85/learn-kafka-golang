package cmd

import (
	"github.com/mmkader85/learn_kafka_golang/kafka"
	"github.com/spf13/cobra"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Start Kafka consumer",
	Run: func(cmd *cobra.Command, args []string) {
		kafka.StartConsumer()
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)
}
