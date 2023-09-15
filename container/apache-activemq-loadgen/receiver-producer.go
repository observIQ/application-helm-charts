package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/Azure/go-amqp"
)

const host = "apache-activemq"
const topic = "topic://test-topic"
const queue = "queue://test-queue"
const port = "5672"
const username = "admin"
const password = "admin"

// Helper for errors
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %s", msg, err)
	}
}

func startQueueProducer(wg *sync.WaitGroup) {
	defer wg.Done()
	ctx := context.Background()

	host_address := fmt.Sprintf("amqp://%s:%s@%s:%s", username, password, host, port)

	// create connection
	conn, err := amqp.Dial(ctx, host_address, nil)
	failOnError(err, "Connection to server failed")
	defer conn.Close()

	// open a session
	session, err := conn.NewSession(ctx, nil)
	failOnError(err, "Failed to open a session.")

	// Send a message
	{
		var count = 100
		// Create a sender
		sender, err := session.NewSender(ctx, queue, nil)
		failOnError(err, "Failed to create new sender")
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

		for i := 0; i < count; i++ {
			// Send message
			err = sender.Send(ctx, amqp.NewMessage([]byte("Hello World!")), &amqp.SendOptions{})
			failOnError(err, "Failed to send message")
		}

		sender.Close(ctx)
		cancel()
	}
}

func startTopicProducer(wg *sync.WaitGroup) {
	defer wg.Done()
	ctx := context.Background()

	host_address := fmt.Sprintf("amqp://%s:%s@%s:%s", username, password, host, port)

	// create connection
	conn, err := amqp.Dial(ctx, host_address, nil)
	failOnError(err, "Connection to server failed")
	defer conn.Close()

	// open a session
	session, err := conn.NewSession(ctx, nil)
	failOnError(err, "Failed to open a session.")

	// Send a message
	{
		var count = 100
		// Create a sender
		sender, err := session.NewSender(ctx, topic, nil)
		failOnError(err, "Failed to create new sender")
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

		for i := 0; i < count; i++ {
			// Send message
			err = sender.Send(ctx, amqp.NewMessage([]byte("Hello World!")), &amqp.SendOptions{})
			failOnError(err, "Failed to send message")
		}

		sender.Close(ctx)
		cancel()
	}
}

func startQueueReceiver(wg *sync.WaitGroup) {
	defer wg.Done()
	ctx := context.Background()

	host_address := fmt.Sprintf("amqp://%s:%s@%s:%s", username, password, host, port)

	// create connection
	conn, err := amqp.Dial(ctx, host_address, nil)
	failOnError(err, "Connection to server failed")
	defer conn.Close()

	// open a session
	session, err := conn.NewSession(ctx, nil)
	failOnError(err, "Failed to open a session.")

	// continuously read messages from a queue
	{
		// Create a receiver
		receiver, err := session.NewReceiver(ctx, queue, &amqp.ReceiverOptions{})
		failOnError(err, "Failed creating receiver link")
		defer func() {
			ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
			receiver.Close(ctx)
			cancel()
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

		for {
			// receive next message
			msg, err := receiver.Receive(ctx, nil)
			if err != nil {
				log.Fatal("Reading message from AMQP:", err)
			}

			// accept message
			if err = receiver.AcceptMessage(context.Background(), msg); err != nil {
				log.Fatalf("Failure accepting message: %v", err)
			}

			fmt.Printf("Message received: %s\n", msg.GetData())
		}
	}
}

func startTopicReceiver(wg *sync.WaitGroup) {
	defer wg.Done()
	ctx := context.Background()

	host_address := fmt.Sprintf("amqp://%s:%s@%s:%s", username, password, host, port)

	// create connection
	conn, err := amqp.Dial(ctx, host_address, nil)
	failOnError(err, "Connection to server failed")
	defer conn.Close()

	// open a session
	session, err := conn.NewSession(ctx, nil)
	failOnError(err, "Failed to open a session.")

	// continuously read messages from a topic
	{
		// Create a receiver
		receiver, err := session.NewReceiver(ctx, topic, &amqp.ReceiverOptions{})
		failOnError(err, "Failed creating receiver link")
		defer func() {
			ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
			receiver.Close(ctx)
			cancel()
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

		for {
			// receive next message
			msg, err := receiver.Receive(ctx, nil)
			if err != nil {
				log.Fatal("Reading message from AMQP:", err)
			}

			// accept message
			if err = receiver.AcceptMessage(context.Background(), msg); err != nil {
				log.Fatalf("Failure accepting message: %v", err)
			}

			fmt.Printf("Message received: %s\n", msg.GetData())
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go startTopicReceiver(&wg)
	go startTopicProducer(&wg)
	go startQueueReceiver(&wg)
	go startQueueProducer(&wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
