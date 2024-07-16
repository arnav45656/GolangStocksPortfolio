package utils

import (
	"encoding/json"
	"log"
	"path"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/gin-gonic/gin"
)

func RabbitMQHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)
		status := c.Writer.Status()
		ClientIP := c.ClientIP()
		method := c.Request.Method
		path := path.Clean(c.Request.URL.Path)
		userAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		datalenght := c.Writer.Size()
		if datalenght < 0 {
			datalenght = 0
		}

		logMessage := map[string]interface{}{
			"status":     status,
			"duration":   duration,
			"ClientIP":   ClientIP,
			"method":     method,
			"path":       path,
			"userAgent":  userAgent,
			"referer":    referer,
			"datalenght": datalenght,
			"timestamp":  time.Now().Format(time.RFC3339),
		}

		SendToRabbitMQ(logMessage)

	}
}

func SendToRabbitMQ(logMessage map[string]interface{}) {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"logs", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	body, err := json.Marshal(logMessage)
	if err != nil {
		log.Fatalf("Failed to marshal log message: %v", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	log.Printf(" [x] Sent %s\n", body)

}
