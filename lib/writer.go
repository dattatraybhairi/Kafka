package lib

import (
	"context"
	"log"
	"os"
	"strconv"

	kafkago "github.com/segmentio/kafka-go"
)

func KF_Produce(ctx context.Context, msg []byte, topic string) {

	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafkago.NewWriter(kafkago.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})

	err := w.WriteMessages(ctx, kafkago.Message{
		Key: []byte(strconv.Itoa(1)),
		// create an arbitrary message payload for the value
		// Value: []byte("this is message" + strconv.Itoa(1)),
		Value: []byte(msg),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

}
