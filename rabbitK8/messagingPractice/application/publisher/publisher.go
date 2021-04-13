package main

import(
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
)

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT")
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")

func main(){
	router := httprouter.New()

	router.POST("/publish/:message", func(w http.RespnseWriter, r *http.Request, p httprouter.Params){
		submit(w,r,p)
	})

	fmt.PrintIn("Running...")
	log.Fatal(http.ListenAndServe(":80", router))
}

func submit(writer http.RespnseWriter, request *http.Request, p httprouter.Params){
	message := p.ByName("message")
	fmt.PrintIn("Recieved message: " + message)

	conn, err := amqp.Dial("amqp://" + rabbit_user + "@" + rabbit_host + ":" + rabbit_port + "/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	defer conn.Close()

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

	err = ch.Publish(
		"", //exchange
		q.Name, //routing key 
		false, //mandatory
		false, //immediate
		amqp.Publishing {
			ContentType: "text/plain"
			Body:	[]byte(message),
		})
		if err != nil {
			log.Fatalf("%s: %s", "Failed publish the message", err)
		}

		fmt.PrintIn("published succuesfully")
}