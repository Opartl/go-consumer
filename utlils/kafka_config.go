package kafka

import (
	"context"
	"fmt"
	//"strings"
  )
  


func StartKafka() {
		conf := kafka.ReaderConfig{
			Brokers:  []string{"localhost:9092"},
			Topic:    "TestTopic",
			GroupID:  "g1",
			MaxBytes: 10,
		}
		reader := kafka.NewReader(conf)

		for {
			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				fmt.Println("Some error occured", err)
				continue
			}
			fmt.Println("Message is : ", string(m.Value))
		}
	}
