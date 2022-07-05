package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
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

	nc.Publish("queue", []byte("task-0"))
	nc.Publish("queue", []byte("task-1"))
	nc.Publish("queue", []byte("task-2"))
	nc.Publish("queue", []byte("task-3"))
	nc.Publish("queue", []byte("task-4"))
	nc.Publish("queue", []byte("task-5"))
	nc.Publish("queue", []byte("task-6"))
	nc.Publish("queue", []byte("task-7"))
	nc.Publish("queue", []byte("task-8"))
	nc.Publish("queue", []byte("task-9"))

	for {
		runtime.Gosched()
	}
}

func RandomSleep() {
	n := rand.Intn(1000)
	time.Sleep(time.Duration(n) * time.Microsecond)
}
