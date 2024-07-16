package adapters

import (
	"api_crud/core/setting"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQChannelQueue struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQChannelQueue(conn *amqp.Connection, queueName string) (*RabbitMQChannelQueue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	//defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &RabbitMQChannelQueue{
		channel: ch,
		queue:   q,
	}, nil
}

func NewRabbitConnection(rabbitSetting setting.RabbitMQ) (*amqp.Connection, error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s", rabbitSetting.User, rabbitSetting.Password, rabbitSetting.Host)

	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}
	//defer func(conm *amqp.Connection) {
	//	err := conm.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(conm)

	return conn, nil
}
