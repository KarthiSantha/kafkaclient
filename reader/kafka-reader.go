package reader

import (
    "fmt"
	"github.com/segmentio/kafka-go"
	"context"
	"kafkaclient/Constants"
)

func Reader(){
	fmt.Printf("Kafka Reader started to read ---------  \n")
		// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   Constants.BootstrapServers,
		Topic:     Constants.Topic,
		GroupID:   Constants.GroupID,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
	
	r.Close()
}