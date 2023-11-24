package main

import (
	"net/http"
	"fmt"
	"encoding/json"
  
	"github.com/gin-gonic/gin"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type LogEntry struct {
	Level       string            `json:"level"`
	Message     string            `json:"message"`
	ResourceID  string            `json:"resourceId"`
	Timestamp   string            `json:"timestamp"`
	TraceID     string            `json:"traceId"`
	SpanID      string            `json:"spanId"`
	Commit      string            `json:"commit"`
	Metadata    map[string]string `json:"metadata"`
}

func main() {
	r := gin.Default()

	topic := "logs"
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id": "logger",
		"acks": "all",
	})
	
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	fmt.Printf("%+v\n", p)

	r.POST("/", func(c *gin.Context) {
		var logEntry LogEntry
		if err := c.ShouldBindJSON(&logEntry); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		level := logEntry.Level
		message := logEntry.Message
		resourceID := logEntry.ResourceID

		fmt.Printf("\nLevel: %s Message: %s ResourceID: %s", level, message, resourceID)

		jsonData, err := json.Marshal(logEntry)

		if err != nil{
			fmt.Printf("%s", "Unable to serialize json to byte array")
		}

		delivery_chan := make(chan kafka.Event, 10000)
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value: []byte(jsonData)},
			delivery_chan,
		)

		if err != nil {
			fmt.Printf("%s", "Unable to produce message")
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	})

	// r.GET("/ping", func(c *gin.Context) {
	//   c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong",
	//   })
	// })

	r.Run(":3000")
}