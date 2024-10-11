package mq

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)


func InitRabbit() (*amqp.Connection, *amqp.Channel, *amqp.Queue, error){
    mq_url := fmt.Sprintf(
        "amqp://%s:%s@%s:%s/",
        rabbit_username, rabbit_password, rabbit_address, rabbit_port)
    fmt.Printf("rabbitmq url: %v\n", mq_url)
    conn, err := amqp.Dial(mq_url)
    if err != nil {
        log.Printf("failed to connect rabbitmq: %v\n", err)
        return nil, nil, nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        log.Printf("failed to open rabbitmq channel: %v\n", err)
        conn.Close()
        return nil, nil, nil, err
    }
    
    q, err := ch.QueueDeclare(
        rabbit_queue, // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    if err != nil {
        log.Printf("failed to declare rabbitmq queue: %v\n", err)
        ch.Close()
        conn.Close()
        return nil, nil, nil, err
    }

    return conn, ch, &q, nil
}

func RabbitSend(conn *amqp.Connection, ch *amqp.Channel, q *amqp.Queue, msg string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    defer conn.Close()
    defer ch.Close()

    err := ch.PublishWithContext(ctx,
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            ContentType: "text/plain",
            Body:        []byte(msg),
        })
    if err != nil {
        log.Printf("faieldt to send message to rabbitmq: %v\n", err)
        return err
    }
    return nil
}


func RabbitReceive(conn *amqp.Connection, ch *amqp.Channel, q *amqp.Queue) (string, error) {
    defer conn.Close()
    defer ch.Close()

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        log.Printf("faieldt to consume message from rabbitmq: %v\n", err)
        return "", err
    }
    
    msg := <-msgs
    return string(msg.Body), nil
}
