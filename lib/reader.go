package lib

import (
	"context"
	"fmt"
	"log"
	"os"

	kafkago "github.com/segmentio/kafka-go"
)

const (
	topic         = "mytopicNew"
	brokerAddress = "34.123.54.243:9093"
)

func KF_Consume(ctx context.Context, topic string) {

	fmt.Println("connecting to the broker:", brokerAddress)
	fmt.Println("Inside Consumer :  consume", topic)

	l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers:   []string{brokerAddress},
		Topic:     topic,
		Partition: 0,
		// GroupID: "my-group",
		// assign the logger to the reader
		Logger: l,
	})

	// r1 := kafkago.NewReader(kafkago.ReaderConfig{
	// 	Brokers: []string{brokerAddress},
	// 	Topic:   "datta",
	// 	// GroupID: "my-group",
	// 	// assign the logger to the reader
	// 	Logger: l,
	// })

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value), "| topic : ", string(msg.Topic), "| Partition: ", string(msg.Partition))
		// r.Close()

		// msg1, err1 := r1.ReadMessage(ctx)
		// if err1 != nil {
		// 	fmt.Println("could not read message " + err1.Error())
		// }
		// // after receiving the message, log its value
		// fmt.Println("received: ", string(msg1.Value))
		// r1.Close()
	}
}
