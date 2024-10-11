package handler

import (
	"net/http"
	"start-backend/mq"

	"github.com/gin-gonic/gin"
)

func KafkaProduce(c *gin.Context) {
    conn := mq.KafkaDb()
    message := c.Params.ByName("message")
    err := mq.KafkaProduce(conn, message)
    if err != nil {
        c.JSON(
            http.StatusInternalServerError,
            gin.H{"failed to produce to kafka": err.Error()})
    }
    c.JSON(http.StatusOK, gin.H{"message": "Message sent to Kafka"})
}

func KafkaConsume(c *gin.Context) {
    conn := mq.KafkaDb()
    message, err := mq.KafkaConsume(conn)
    if err != nil {
        c.JSON(
            http.StatusInternalServerError,
            gin.H{"failed to consume message from kafka": err.Error()})
    }
    c.JSON(http.StatusOK, gin.H{"message": message})
}
