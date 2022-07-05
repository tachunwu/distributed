package main

import (
	"fmt"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestQueueGroupPub(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// Worker pool a: worker-0
	nc.QueueSubscribe("queue", "worker_pool_a", func(m *nats.Msg) {
		fmt.Println("worker_pool_a: worker-0", string(m.Data))
	})
	// Worker pool a: worker-1
	nc.QueueSubscribe("queue", "worker_pool_a", func(m *nats.Msg) {
		fmt.Println("worker_pool_a: worker-1", string(m.Data))
	})
	// Worker pool a: worker-2
	nc.QueueSubscribe("queue", "worker_pool_a", func(m *nats.Msg) {
		fmt.Println("worker_pool_a: worker-2", string(m.Data))
	})

	// Worker pool b: worker-0
	nc.QueueSubscribe("queue", "worker_pool_b", func(m *nats.Msg) {
		fmt.Println("worker_pool_b: worker-0", string(m.Data))
	})
	// Worker pool b: worker-1
	nc.QueueSubscribe("queue", "worker_pool_b", func(m *nats.Msg) {
		fmt.Println("worker_pool_b: worker-1", string(m.Data))
	})
	// Worker pool b: worker-2
	nc.QueueSubscribe("queue", "worker_pool_b", func(m *nats.Msg) {
		fmt.Println("worker_pool_b: worker-2", string(m.Data))
	})

	nc.Publish("queue", []byte("Pub task-0 to queue"))
	nc.Publish("queue", []byte("Pub task-1 to queue"))
	nc.Publish("queue", []byte("Pub task-2 to queue"))
	nc.Publish("queue", []byte("Pub task-3 to queue"))
	nc.Publish("queue", []byte("Pub task-4 to queue"))
	nc.Publish("queue", []byte("Pub task-5 to queue"))
	nc.Publish("queue", []byte("Pub task-6 to queue"))
	nc.Publish("queue", []byte("Pub task-7 to queue"))
	nc.Publish("queue", []byte("Pub task-8 to queue"))
	nc.Publish("queue", []byte("Pub task-9 to queue"))

}
