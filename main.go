package main

import (
  "fmt"
	kafka "kafkalocal/utils"
	"time"
)

func main() {

	fmt.Println("Starting")
	go kafka.StartKafka()
	fmt.Println("Kafka has been started...")
	time.Sleep(10 * time.Minute)

}
