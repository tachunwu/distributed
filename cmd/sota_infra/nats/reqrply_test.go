package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
)

func TestReqRply(t *testing.T) {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	nc.Subscribe("req", func(m *nats.Msg) {
		RandomSleep()
		nc.Publish(m.Reply, []byte("rply from subscriber-0"))
	})

	nc.Subscribe("req", func(m *nats.Msg) {
		RandomSleep()
		nc.Publish(m.Reply, []byte("rply from subscriber-1"))
	})

	nc.Subscribe("req", func(m *nats.Msg) {
		RandomSleep()
		nc.Publish(m.Reply, []byte("rply from subscriber-2"))
	})

	msg, err := nc.Request("req", []byte("req"), 10*time.Millisecond)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(msg.Data))
	}

	for {
		runtime.Gosched()
	}
}
