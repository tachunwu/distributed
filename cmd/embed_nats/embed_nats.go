package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

func main() {
	js := JetStreamInit()
	SubscribeSingleHomeTxn(js)
	SubscribeMultiHomeTxn(js)

	for {
		runtime.Gosched()
	}

}

func JetStreamInit() nats.JetStreamContext {
	opts := &server.Options{
		JetStream: true,
		StoreDir:  "./data",
		Cluster:   server.ClusterOpts{},
	}

	// Initialize new server with options
	ns, err := server.NewServer(opts)

	if err != nil {
		panic(err)
	}

	// Start the server via goroutine
	go ns.Start()

	// Wait for server to be ready for connections
	if !ns.ReadyForConnections(4 * time.Second) {
		panic("not ready for connection")
	}

	// Connect to server
	nc, err := nats.Connect(ns.ClientURL())
	js, _ := nc.JetStream()

	if err != nil {
		panic(err)
	}
	// Create a Stream
	js.AddStream(&nats.StreamConfig{
		Name:     "SINGLE",
		Subjects: []string{"SLOG.single"},
	})
	js.AddStream(&nats.StreamConfig{
		Name:     "MULTI",
		Subjects: []string{"SLOG.multi"},
	})

	// Create a Consumer
	js.AddConsumer("SLOG", &nats.ConsumerConfig{
		Durable: "SCHED",
	})

	return js
}

func SubscribeMultiHomeTxn(js nats.JetStreamContext) {
	subMulti, _ := js.PullSubscribe("SLOG.multi", "SCHED")
	for {
		msgs, _ := subMulti.Fetch(128)
		for _, msg := range msgs {
			fmt.Println(msg)
		}
	}
}

func SubscribeSingleHomeTxn(js nats.JetStreamContext) {
	subSingle, _ := js.PullSubscribe("SLOG.single", "SCHED")
	for {
		msgs, _ := subSingle.Fetch(128)
		for _, msg := range msgs {
			fmt.Println(msg)
		}
	}
}

func PublishTestTxns(js nats.JetStreamContext) {
	// Simple Async Stream Publisher
	for i := 0; i < 1000000; i++ {
		js.PublishAsync("SLOG.single", []byte("single"))
	}
	for i := 0; i < 1000000; i++ {
		js.PublishAsync("SLOG.multi", []byte("multi"))
	}
}
