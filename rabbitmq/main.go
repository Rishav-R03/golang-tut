package main

/**
Key Concepts to Remember
The Connection vs. The Channel:
	Think of the Connection as the physical fiber-optic cable running to a house,
	and the Channel as an individual phone call being made over that cable. You open one connection but can have
	many channels for different tasks.

The Queue Declaration:
	In RabbitMQ, "Declare" is an idempotent operation.
	It means "Make sure this exists.
	" This is safer than a simple "Create" command
	because it won't throw an error if the queue is already there.

Networking in Docker:
	Since your config.yml uses the URL amqp://guest:guest@rabbitmq:5672/,
	your Go app will look for a host named rabbitmq. This works because Docker Compose creates a DNS entry for each service name in its internal network.

*/
import (
	"log"
	"time"

	// Using the official community RabbitMQ client driver
	"rabbitmq/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// 1. Load Configuration
	// WHY: We separate infrastructure details (URLs, ports) from logic so we can
	// change environments (Dev vs. Prod) without recompiling the code.

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading: %v", err)
	}

	// 2. Establish a Connection
	// WHY: This is a long-lived TCP connection between your Go app and the RabbitMQ server.
	// It handles authentication and heartbeat frames to keep the link alive.

	conn, err := amqp.Dial(cfg.RabbitMQ.URL)
	if err != nil {
		// WHY: If the broker isn't up, the app cannot function. Fatalf logs and exits the program.
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	// 3. Defer Connection Closure
	// WHY: 'defer' ensures the connection closes gracefully when main() finishes,
	// preventing "zombie" connections on the RabbitMQ server side.

	defer conn.Close()

	// 4. Open a Channel
	// WHY: TCP connections are expensive to open/close. RabbitMQ uses "Channels"
	// which are virtual connections inside the TCP pipe. Most API work happens here.
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	// 5. Defer Channel Closure
	// WHY: Similar to the connection, we want to clean up our virtual pipe
	// as soon as this function scope ends.
	defer ch.Close()

	// 6. Declare a Queue
	// WHY: RabbitMQ is flexible. You can publish to non-existent queues, but declaring
	// ensures the queue exists with the properties you expect (durable, exclusive, etc.)

	q, err := ch.QueueDeclare(
		cfg.RabbitMQ.QueueName, // name of queue
		false,                  // durable (survives server restarts?)
		false,                  //delete when unused?
		false,                  // exclusive (only used by this connection?)
		false,                  //no wait?
		nil,                    //arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v ", err)
	}

	// 7. Success check
	log.Printf("Successfully connected to RabbitMQ and declared queue: %s", q.Name)

	// error because we haven't published or consumed any messages yet, but the connection and queue setup is successful.

	// 8. Create a messsage
	body := "Hello RabbitMQ!"

	//9. Start a loop to send msgs to the queue every 5 seconds
	// WHY: This simulates a producer that continuously sends data to the queue.
	// In a real app, this could be triggered by user actions, scheduled tasks, etc.

	for {
		err = ch.Publish(
			"",     // exchange (using default)
			q.Name, // routing key
			false,  //mandatory
			false,  //immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			log.Printf("Failed to publish a message: %v", err)
		} else {
			log.Printf("Sent message: %s", body)
		}
		// Sleep for 5 seconds before sending the next message
		time.Sleep(5 * time.Second)
	}
}
