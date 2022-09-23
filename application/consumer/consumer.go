package main


import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
)

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT") 
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")

func main() {
	consume()
}

func consume() {

	// Establish connection ke rabbit host

	conn, err := amqp.Dial("amqp://" + rabbit_user + ":" +rabbit_password + "@" + rabbit_host + ":" + rabbit_port +"/")

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	// Establish channel menuju queue

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	// Declare Exchange untuk menyambung koneksi dengan exhange dari publisher

	err = ch.ExchangeDeclare(
		"posts",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare an exchange", err)
	}

	// Declare queue 

	q, err := ch.QueueDeclare(
		"task", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	// Untuk bind ke message queue

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"posts", // exchange
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to bind a queue", err)
	}


	fmt.Println("Channel and Queue established")

	defer conn.Close()
	defer ch.Close()

	// Menangkap message dari queue

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	  )

	if err != nil {
		log.Fatalf("%s: %s", "Failed to register consumer", err)
	}

	// Agar message yang ditangkap bisa terus-terusan

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			
			d.Ack(false)
		}
	  }()
	  
	  fmt.Println("Running...")
	  <-forever
}