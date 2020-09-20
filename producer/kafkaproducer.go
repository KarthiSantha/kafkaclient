package Producer

import (
    "fmt"
	"github.com/segmentio/kafka-go"
	"context"
	"kafkaclient/Constants"
	"io/ioutil"
)


func Producer(){

	// make a writer that produces to topic-A, using the least-bytes distribution
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: Constants.BootstrapServers,
		Topic:   Constants.Topic,
		Balancer: &kafka.LeastBytes{},
	})

	// Read file and data json
	file, _ := ioutil.ReadFile("test.json")

	fmt.Printf("\n Kafka Producer is created with Bootstarp Server : %v and Topic %v \n",Constants.BootstrapServers,Constants.Topic)
	w.WriteMessages(context.Background(),
		kafka.Message{Value: []byte(file)},
	)

	w.Close()


}

