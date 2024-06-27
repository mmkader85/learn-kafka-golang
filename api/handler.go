package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmkader85/learn_kafka_golang/kafka"
)

func SendMessage(c *gin.Context) {
	var json map[string]interface{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := json.Marshal(json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	kafka.ProduceMessage(message)
	c.JSON(http.StatusOK, gin.H{"status": "message sent to Kafka"})
}
