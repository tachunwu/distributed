package main

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestPubSub(t *testing.T) {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	nc.Subscribe("pubsub", func(m *nats.Msg) {
		fmt.Println("Sub from: pubsub", string(m.Data))
	})

	nc.Subscribe("pubsub.*", func(m *nats.Msg) {
		fmt.Println("Sub from: pubsub.*", string(m.Data))
	})

	nc.Subscribe("pubsub.>", func(m *nats.Msg) {
		fmt.Println("Sub from: pubsub.>", string(m.Data))
	})

	// Pub-Sub
	nc.Publish("pubsub", []byte("Pub to pubsub"))
	nc.Publish("pubsub.0", []byte("Pub to pubsub.0"))
	nc.Publish("pubsub.1", []byte("Pub to pubsub.1"))
	nc.Publish("pubsub.2", []byte("Pub to pubsub.2"))
	nc.Publish("pubsub.0.0", []byte("Pub to pubsub.0.0"))
	nc.Publish("pubsub.0.1", []byte("Pub to pubsub.0.1"))
	nc.Publish("pubsub.0.2", []byte("Pub to pubsub.0.2"))

	for {
		runtime.Gosched()
	}
}
