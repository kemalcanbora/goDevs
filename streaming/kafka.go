package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	helper "goDevs/helper"
	"log"
	"os"
	"time"
)

func Producer(topic string, data []byte, key []byte) {
	helper.GetEnv()
	kafkaHost := os.Getenv("DOCKER_KAFKA_HOST")
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaHost, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	_, err = conn.WriteMessages(
		kafka.Message{Value: data, Key: key},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
