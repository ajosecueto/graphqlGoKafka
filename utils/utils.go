package utils

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"graphqlGoKafka/graph/model"
	"log"
	"strings"
	"time"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  1 * time.Second,
	})
}

func StartConsumer(ctx context.Context) {
	kafkaURL := "localhost:9092"
	topic := "kafka2"
	groupID := "groupID"

	reader := getKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		if m.Topic == "create_profile" {
			// check Node based in user_id
			// Create Node
		}

		if m.Topic == "edit_profile" {
			// update Node where user_id
		}

		if m.Topic == "delete_profile" {
			// delete Node where user_id and relationship
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}

func SendPreferenceEvent(input model.NewPreference) {
	topic := "kafka2"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(input.Description)},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
