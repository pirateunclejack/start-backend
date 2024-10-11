package mq

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func KafkaDb() *kafka.Conn{
    conn, err := kafka.DialLeader(
        context.Background(),
        "tcp",
        kafka_address,
        kafka_topic,
        kafka_partition,
    )
    if err != nil {
        log.Printf("failed to connect to kafka: %v\n", err)
        return nil
    }

    return conn
}

func KafkaReader() *kafka.Reader {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{kafka_address},
        Topic:   kafka_topic,
        Partition: kafka_partition,
        MaxBytes: 10e3,
        MaxWait: 3*time.Second,
    })

    return r
}


func KafkaProduce(conn *kafka.Conn, message string) error {
    conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
    defer conn.Close()
    _, err := conn.WriteMessages(
        kafka.Message{Value: []byte(message)},
    )
    if err != nil {
        log.Printf("failed to create message to kafka: %v\n", err)
        return err
    }

    return nil
}

func KafkaConsume(conn *kafka.Conn) (string, error) {
    defer conn.Close()
    reader := KafkaReader()
    offset, err := conn.ReadLastOffset()

    log.Println("offset:", offset)
    if err != nil {
        log.Printf("failed to get offset: %v\n", err)
        return "", nil
    }
    reader.SetOffset(offset-1)
    msg, err := reader.ReadMessage(context.Background())
    if err != nil {
        log.Printf("failed to consume message: %v\n", err)
        return "", err
    }
    if err := reader.Close(); err != nil {
        log.Printf("failed to close reader:%v\n", err)
        return "", err
    }
    return string(msg.Value), nil
}
