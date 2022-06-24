package main

func main() {

}

type CalvinFactory interface {
	createSequencer() Sequencer
	createScheduler() Scheduler
	createStorager() Storager
}

func NewCalvinFactory() (CalvinFactory, error) {
	return nil, nil
}

type Sequencer interface {
}

type Scheduler interface {
}

type Storager interface {
	Get()
	Set()
	Delete()
	Scan()
}

type Storage struct {
	path string
}

type PebbleStorage struct {
	Storage
}

type BadgerStorage struct {
	Storage
}
