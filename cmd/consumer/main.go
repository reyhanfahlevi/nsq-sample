package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/reyhanfahlevi/pkg/go/nsq"
)

func main() {
	c := nsq.NewConsumer(nsq.ConsumerConfig{
		ListenAddress: []string{"localhost:4161"},
		Prefix:        "",
	})

	c.RegisterHandler(nsq.ConsumerHandler{
		Topic:       "test",
		Channel:     "rey",
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

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case <-term:
		log.Println("application terminated ")
	}
}

func handlerSample(message nsq.IMessage) error {
	msg := struct {
		Test string `json:"test"`
	}{
		Test: "test",
	}

	err := json.Unmarshal(message.GetBody(), &msg)
	if err != nil {
		return err
	}

	fmt.Println(msg)

	return nil
}
