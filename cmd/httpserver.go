package cmd

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mmkader85/learn-kafka-golang/api"
	"github.com/spf13/cobra"
)

var httpServerCmd = &cobra.Command{
	Use:   "httpserver",
	Short: "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()
		r := gin.Default()
		r.POST("/send", api.SendMessage)
		r.Run(":" + os.Getenv("HTTP_SERVER_PORT"))
	},
}

func init() {
	rootCmd.AddCommand(httpServerCmd)
}
