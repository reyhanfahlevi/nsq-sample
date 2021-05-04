package main

import (
	"log"

	"github.com/reyhanfahlevi/pkg/go/nsq"
)

func main() {
	pb, err := nsq.NewPublisher("127.0.0.1:4150", "")
	if err != nil {
		log.Fatal(err)
	}

	msg := struct {
		Test string `json:"test"`
	}{
		Test: "Coba Test",
	}
	_ = pb.PublishWithoutPrefix("test", msg)
}
