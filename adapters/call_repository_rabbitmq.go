package adapters

import (
	"api_crud/domain"
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type CallRabbitMQRepository struct {
	Conn        *amqp.Connection
	QueueChanel *RabbitMQChannelQueue
}

func NewCallRabbitMQRepository(conn *amqp.Connection) *CallRabbitMQRepository {
	qc, err := NewRabbitMQChannelQueue(conn, "call_create")
	if err != nil {
		fmt.Printf("Error creating Rabbit")
		return nil
	}
	return &CallRabbitMQRepository{
		Conn:        conn,
		QueueChanel: qc,
	}
}

func (c CallRabbitMQRepository) SendCall(ctx context.Context, call *domain.Call) error {
	body, err := json.Marshal(call)
	if err != nil {
		log.Fatalln("Failed to marshal struct: %v", err)
	}
	err = c.QueueChanel.channel.PublishWithContext(ctx,
		"",                       // exchange
		c.QueueChanel.queue.Name, // routing key
		false,                    // mandatory
		false,                    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	fmt.Println("send ")
	//c.channel.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c CallRabbitMQRepository) GetCallAfterCreated(handleUpdate func(call domain.Call)) {

	var forever chan struct{}

	msgs, err := c.QueueChanel.channel.Consume(
		c.QueueChanel.queue.Name, // queue
		"",                       // consumer
		true,                     // auto-ack
		false,                    // exclusive
		false,                    // no-local
		false,                    // no-wait
		nil,                      // args
	)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for d := range msgs {
			var callCreated domain.Call
			err := json.Unmarshal(d.Body, &callCreated)
			if err != nil {
				log.Printf("Failed to unmarshal JSON: %v", err)
				continue
			}
			fmt.Println("Received a data")
			handleUpdate(callCreated)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
