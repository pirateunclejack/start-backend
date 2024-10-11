package handler

import (
	"net/http"
	"start-backend/mq"

	"github.com/gin-gonic/gin"
)

func RabbitSend(c *gin.Context)  {
    conn, ch, q, err := mq.InitRabbit()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to rabbit"})
        return
    }

    message := c.Params.ByName("message")
    err = mq.RabbitSend(conn, ch, q, message)
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "send message to rabbit"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "success to send message to rabbitmq"})
}

func RabbitReceive(c *gin.Context){
    conn, ch, q, err := mq.InitRabbit()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to rabbit"})
        return
    }
    message, err  := mq.RabbitReceive(conn, ch, q)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to consume message from rabbit"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": message})

}
