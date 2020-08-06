package Producer

import (
    "fmt"
	"github.com/segmentio/kafka-go"
	"context"
	"kafkaclient/Constants"
)


func Producer(){

	// make a writer that produces to topic-A, using the least-bytes distribution
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: Constants.BootstrapServers,
		Topic:   Constants.Topic,
		Balancer: &kafka.LeastBytes{},
	})

	fmt.Printf("\n Kafka Producer is created with Bootstarp Server : %v and Topic %v \n",Constants.BootstrapServers,Constants.Topic)
	w.WriteMessages(context.Background(),
		kafka.Message{Value: []byte("Hello World!")},
		kafka.Message{Value: []byte("One!")},
		kafka.Message{Value: []byte("Two!")},
	)

	w.Close()


}

