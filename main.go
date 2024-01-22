package main

import (
    "context"
    "fmt"
    "log"

    "github.com/segmentio/kafka-go"
)

func main() {
    // Configure Kafka consumer
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "my-kafka-topic",
        GroupID: "my-consumer-group", // Optional consumer group ID
    })

    defer r.Close()

    // Consume messages in a loop
    for {
        msg, err := r.ReadMessage(context.Background())
        if err != nil {
            log.Fatalln(err)
        }

        fmt.Printf("Received message: %s\n", string(msg.Value))
        // Process the message as needed
    }
}
