package main

import (
	"log"

	"github.com/reyhanfahlevi/pkg/go/nsq"
)

func main() {
	pb, err := nsq.NewPublisher("localhost:4151", "")
	if err != nil {
		log.Fatal(err)
	}

	msg := struct {
		Test string `json:"test"`
	}{
		Test: "test",
	}
	_ = pb.PublishWithoutPrefix("test", msg)
}
