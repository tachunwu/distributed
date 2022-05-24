package seq

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	calvinpb "github.com/tachunwu/distributed/pkg/proto/calvin/v1"
	"google.golang.org/protobuf/proto"
)

var natsCluster string = "localhost:4000"
var partitionId string = "N0"
var batchSize int = 256

type Sequencer struct {
	calvinpb.UnimplementedSequencerServiceServer
	nc     *nats.Conn
	stream nats.JetStreamContext
}

func NewSequencer() *Sequencer {
	nc, err := nats.Connect(natsCluster)
	if err != nil {
		log.Println(err)
	}

	s, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Println(err)
	}

	// Create stream
	s.AddStream(&nats.StreamConfig{
		Name:       "CALVIN",
		Subjects:   []string{"seq"},
		Storage:    nats.FileStorage,
		Duplicates: 2 * time.Minute,
	})

	// Create consumer
	s.AddConsumer("CALVIN", &nats.ConsumerConfig{
		Durable: "N0",
	})
	return &Sequencer{
		nc:     nc,
		stream: s,
	}
}

func (s *Sequencer) Run() {
	go s.PullSubscribe()
}

func (s *Sequencer) ReceiveNewTxn(ctx context.Context, req *calvinpb.ReceiveNewTxnRequest) (*calvinpb.ReceiveNewTxnResponse, error) {
	txn, err := proto.Marshal(req.Txn)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	s.stream.PublishAsync("seq", txn)
	select {
	case <-s.stream.PublishAsyncComplete():
		return &calvinpb.ReceiveNewTxnResponse{}, nil
	case <-time.After(5 * time.Second):
		log.Println("Enqueue txn timeout")
		return nil, nil
	}
}

func (s *Sequencer) PullSubscribe() {
	sub, _ := s.stream.PullSubscribe("seq", partitionId, nats.PullMaxWaiting(128))
	for {
		txns, _ := sub.Fetch(batchSize)
		for _, txn := range txns {
			// You can process serialize txn here!!
			// Enqueue to determinisic sched
			txn.Ack()
		}
	}
}
