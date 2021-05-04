package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/reyhanfahlevi/pkg/go/nsq"
)

func main() {
	c := nsq.NewConsumer(nsq.ConsumerConfig{
		ListenAddress: []string{"127.0.0.1:4161"},
		Prefix:        "",
	})

	c.RegisterHandler(nsq.ConsumerHandler{
		Topic:       "test",
		Channel:     "consumer-1",
		Concurrent:  0,
		MaxAttempts: 10,
		MaxInFlight: 1,
		Enable:      true,
		Handler:     handlerSample,
	})

	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}

	c.Wait()
}

func handlerSample(message nsq.IMessage) error {
	msg := struct {
		Test string `json:"test"`
	}{}

	err := json.Unmarshal(message.GetBody(), &msg)
	if err != nil {
		return nil
	}

	// doing something here
	fmt.Println(msg)

	return nil
}
