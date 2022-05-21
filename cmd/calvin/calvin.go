package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

var txnCount int = 10

func main() {

	// Connect to NATS
	n0, err := nats.Connect("nats://localhost:4000")
	if err != nil {
		log.Println(err)
		return
	}
	n1, err := nats.Connect("nats://localhost:4001")
	if err != nil {
		log.Println(err)
		return
	}
	n2, err := nats.Connect("nats://localhost:4002")
	if err != nil {
		log.Println(err)
		return
	}

	// Create JetStream Context
	s0, _ := n0.JetStream(nats.PublishAsyncMaxPending(256))
	s1, _ := n1.JetStream(nats.PublishAsyncMaxPending(256))
	s2, _ := n2.JetStream(nats.PublishAsyncMaxPending(256))

	// Create sequencer stream
	s0.AddStream(&nats.StreamConfig{
		Name:       "CALVIN",
		Subjects:   []string{"CALVIN.*"},
		Storage:    nats.FileStorage,
		Duplicates: 2 * time.Minute,
	})

	// Create consumers
	s0.AddConsumer("CALVIN", &nats.ConsumerConfig{
		Durable: "N0",
	})
	s1.AddConsumer("CALVIN", &nats.ConsumerConfig{
		Durable: "N1",
	})
	s2.AddConsumer("CALVIN", &nats.ConsumerConfig{
		Durable: "N2",
	})

	// Publish txn
	go AsyncPublish(txnCount, s0)
	go AsyncPublish(txnCount, s1)
	go AsyncPublish(txnCount, s2)

	// Pull subscribe
	go PullSubscribe(1, "N0", s0)
	go PullSubscribe(1, "N1", s1)
	go PullSubscribe(1, "N2", s2)

	// Sync subscribe
	// go SyncSubscribe("N0", s0)
	// go SyncSubscribe("N1", s1)
	// go SyncSubscribe("N2", s2)

	// Wait
	for {
		runtime.Gosched()
	}

}

func AsyncPublish(txnCount int, s nats.JetStreamContext) {
	for i := 0; i < txnCount; i++ {
		s.PublishAsync("CALVIN.SEQ", []byte("Txn: "+fmt.Sprint(i)))
	}
	select {
	case <-s.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}
}

func PullSubscribe(batchSize int, consumer string, s nats.JetStreamContext) {
	for {
		sub, err := s.PullSubscribe("CALVIN.SEQ", consumer)
		if err != nil {
			log.Println("Pull: ", err)
		}

		msgs, err := sub.Fetch(batchSize)
		if err != nil {
			log.Println("Pull: ", err)
		}

		for i := range msgs {
			fmt.Sprintln(msgs[i].Data)
		}
	}

}

func SyncSubscribe(consumer string, s nats.JetStreamContext) {
	sub, err := s.SubscribeSync("CALVIN.SEQ", nats.Durable(consumer), nats.MaxDeliver(3))
	if err != nil {
		log.Println(err)
	}
	for {
		m, err := sub.NextMsg(5 * time.Second)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Sprintln(m.Data)
		}
	}

}
