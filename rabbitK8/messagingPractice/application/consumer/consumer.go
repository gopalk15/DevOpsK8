package main

import(
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
)

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT")
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")


func main(){
	consume()
}

func consume(){

	conn, err := amqp.Dial("amqp://" + rabbit_user + "@" + rabbit_host + ":" + rabbit_port + "/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	//Create a virtual connection to a queue
	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	defer ch.Close()

	//Declare a queue
	q, err := ch.QueueDeclare(
		"publisher", //name
		false, //durable
		false, //delete when unused
		false, //exclusive
		false, //no-wait
		nil, //arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	fmt.Println("Channel and Queue established")
	defer conn.Close()
	defer ch.Close()
	
}