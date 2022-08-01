package lib

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func CreateTopics(ctx context.Context, topics []string) {

	// fmt.Println(os.Getenv("KAFKA_ENDPOINT"))
	fmt.Println(brokerAddress)
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}

	var controllerConn *kafka.Conn

	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()
	var topicConfigs []kafka.TopicConfig
	for i := 0; i < len(topics); i++ {
		topicConfigs = append(topicConfigs, kafka.TopicConfig{
			Topic:             topics[i],
			NumPartitions:     1,
			ReplicationFactor: 1,
		})
		fmt.Println(topics[i])
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}
