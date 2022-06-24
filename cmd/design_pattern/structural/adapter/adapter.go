package main

type StreamClient struct{}
type Event struct{}

func (c *StreamClient) WriteEvent(e Event) {

}

type StreamServer interface {
	WriteEvent()
}

type Kafka struct{}
type Nats struct{}
