package main

import (
	"net/url"
	"strings"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/spf13/pflag"
)

// 3 partition setup
var partitions = pflag.StringSlice("peers", []string{"localhost:40000", "loclahost:40001", "localhost:40002"}, "IP address of each partition in the cluster")
var multihome = pflag.String("multi-home", "localhost:50000", "Multi-home ordering stream")
var dataDir = pflag.String("data-dir", "", "Directory to store persistent data")
var partition = pflag.String("partition", "", "Partition id for this partition")

func main() {
	// Init Sequencer
	pflag.Parse()
	// Start micro benchmark

}

type Sequencer struct {
	Partition       string
	Partitions      map[string]nats.JetStreamContext
	MultiHomeStream nats.JetStreamContext
	DataDir         string
	GlobalLog       *nats.Subscription
}

func (s *Sequencer) Start() {

}

func NewSequencer() *Sequencer {
	// Embed nats
	opts := &server.Options{
		JetStream: true,
		StoreDir:  *dataDir,
		Routes:    []*url.URL{},
		Cluster: server.ClusterOpts{
			Host: "localhost",
		},
	}

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

	// Client connect
	nc, err := nats.Connect(strings.Join(*partitions, ","))
	js, _ := nc.JetStream()

	js.AddStream(&nats.StreamConfig{
		Name:     *partition,
		Subjects: []string{"SLOG." + *partition},
	})

	js.AddConsumer("SLOG", &nats.ConsumerConfig{
		Durable: "SCHED",
	})
	return &Sequencer{}
}
